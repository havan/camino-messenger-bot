package metadata

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"
)

type Metadata struct {
	RequestID      string                   `json:"request_id"`
	Sender         string                   `json:"sender"`
	Recipient      string                   `json:"recipient"`
	Cheques        []map[string]interface{} `json:"cheques"`
	Timestamps     map[string]int64         `json:"timestamps"` // map of checkpoints to timestamps in unix milliseconds
	NumberOfChunks uint                     `json:"number_of_chunks"`
	ChunkIndex     uint                     `json:"chunk_index"`
}

func (m *Metadata) ExtractMetadata(ctx context.Context) error {
	mdPairs, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		mdPairs, ok = metadata.FromOutgoingContext(ctx)
		if !ok {
			return fmt.Errorf("metadata not found in context")
		}
	}
	return m.FromGrpcMD(mdPairs)
}

func (m *Metadata) FromGrpcMD(mdPairs metadata.MD) error {
	if requestID, found := mdPairs["request_id"]; found {
		m.RequestID = requestID[0]
	}

	if sender, found := mdPairs["sender"]; found {
		m.Sender = sender[0]
	}

	if recipient, found := mdPairs["recipient"]; found {
		m.Recipient = recipient[0]
	}

	if cheques, found := mdPairs["cheques"]; found {
		chequesJSON := strings.Join(cheques, "")
		if err := json.Unmarshal([]byte(chequesJSON), &m.Cheques); err != nil {
			return fmt.Errorf("error unmarshalling cheques: %v", err)
		}
	}

	if timestamps, found := mdPairs["timestamps"]; found {
		timestampsJSON := strings.Join(timestamps, "")
		if err := json.Unmarshal([]byte(timestampsJSON), &m.Timestamps); err != nil {
			return fmt.Errorf("error unmarshalling timestamps: %v", err)
		}
	}
	return nil
}
func (m *Metadata) ToGrpcMD() metadata.MD {
	md := metadata.New(map[string]string{
		"request_id": m.RequestID,
		"sender":     m.Sender,
		"recipient":  m.Recipient,
		"timestamps": func() string {
			timestampsJSON, _ := json.Marshal(m.Timestamps)
			return string(timestampsJSON)
		}(),
		"cheques": func() string {
			chequesJSON, _ := json.Marshal(m.Cheques)
			return string(chequesJSON)
		}(),
	})
	return md
}
func (m *Metadata) Stamp(checkpoint string) {
	if m.Timestamps == nil {
		m.Timestamps = make(map[string]int64)
	}
	idx := len(m.Timestamps) // for analysis' sake, we want to know the order of the checkpoints
	m.Timestamps[fmt.Sprintf("%d-%s", idx, checkpoint)] = time.Now().UnixMilli()
}
func (m *Metadata) StampOn(checkpoint string, t int64) {
	if m.Timestamps == nil {
		m.Timestamps = make(map[string]int64)
	}
	idx := len(m.Timestamps) // for analysis' sake, we want to know the order of the checkpoints
	m.Timestamps[fmt.Sprintf("%d-%s", idx, checkpoint)] = t
}
