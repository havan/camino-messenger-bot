package matrix

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"

	"camino-messenger-bot/config"
	"camino-messenger-bot/internal/messaging"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/crypto/cryptohelper"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"

	_ "github.com/mattn/go-sqlite3"
)

var _ messaging.Messenger = (*messenger)(nil)

var C4TMessage = event.Type{Type: "m.room.c4t-msg", Class: event.MessageEventType}

type client struct {
	*mautrix.Client
	ctx          context.Context
	cancelSync   context.CancelFunc
	syncStopWait sync.WaitGroup
	cryptoHelper *cryptohelper.CryptoHelper
}
type messenger struct {
	msgChannel chan messaging.Message
	cfg        *config.MatrixConfig
	logger     *zap.SugaredLogger
	client     client
}

func NewMessenger(cfg *config.MatrixConfig, logger *zap.SugaredLogger) *messenger {
	c, err := mautrix.NewClient(cfg.Host, "", "")
	if err != nil {
		panic(err)
	}
	return &messenger{
		msgChannel: make(chan messaging.Message),
		cfg:        cfg,
		logger:     logger,
		client:     client{Client: c},
	}
}

func (m *messenger) StartReceiver() error {
	syncer := m.client.Syncer.(*mautrix.DefaultSyncer)
	// TODO: custom message event types have to be registered properly . see also event.TypeMap[event.MessageEventType] = event.MessageEventContent{}
	syncer.OnEventType(event.EventMessage, func(source mautrix.EventSource, evt *event.Event) {
		m.logger.Debug("Received msg",
			zap.String("sender", evt.Sender.String()),
			zap.String("type", evt.Type.String()),
			zap.String("id", evt.ID.String()),
			zap.String("body", evt.Content.AsMessage().Body))
		m.msgChannel <- messaging.Message{
			Metadata: messaging.Metadata{
				Sender: evt.Sender.String(),
				RoomID: evt.RoomID.String(),
			},
			RequestID: evt.ID.String(),
			Body:      evt.Content.AsMessage().Body,
			Type:      messaging.MessageType(evt.Content.AsMessage().MsgType),
		}
	})
	syncer.OnEventType(event.StateMember, func(source mautrix.EventSource, evt *event.Event) {
		if evt.GetStateKey() == m.client.UserID.String() && evt.Content.AsMember().Membership == event.MembershipInvite {
			_, err := m.client.JoinRoomByID(evt.RoomID)
			if err == nil {
				m.logger.Info("Joined room after invite",
					zap.String("room_id", evt.RoomID.String()),
					zap.String("inviter", evt.Sender.String()))
			} else {
				m.logger.Error("Failed to join room after invite",
					zap.String("room_id", evt.RoomID.String()),
					zap.String("inviter", evt.Sender.String()))
			}
		}
	})

	cryptoHelper, err := cryptohelper.NewCryptoHelper(m.client.Client, []byte("meow"), m.cfg.Store) //TODO refactor
	if err != nil {
		return err
	}

	cryptoHelper.LoginAs = &mautrix.ReqLogin{
		Type:       mautrix.AuthTypePassword,
		Identifier: mautrix.UserIdentifier{Type: mautrix.IdentifierTypeThirdParty, Medium: "camino", Address: m.cfg.Username},
		Password:   m.cfg.Password,
	}
	err = cryptoHelper.Init()
	if err != nil {
		return err
	}
	// Set the client crypto helper in order to automatically encrypt outgoing messages
	m.client.Crypto = cryptoHelper
	m.client.cryptoHelper = cryptoHelper // nikos: we need the struct cause stop method is not available on the interface level

	m.logger.Info("Now running")
	syncCtx, cancelSync := context.WithCancel(context.Background())
	m.client.ctx = syncCtx
	m.client.cancelSync = cancelSync
	m.client.syncStopWait.Add(1)

	go func() {
		err = m.client.SyncWithContext(syncCtx)
		defer m.client.syncStopWait.Done()
		if err != nil && !errors.Is(err, context.Canceled) {
			panic(err)
		}
	}()

	return nil
}
func (m *messenger) StopReceiver() error {
	m.logger.Info("Stopping matrix syncer...")
	if m.client.cancelSync != nil {
		m.client.cancelSync()
	}
	m.client.syncStopWait.Wait()
	return m.client.cryptoHelper.Close()
}

func (m *messenger) SendAsync(ctx context.Context, msg messaging.Message) error {
	m.logger.Info("Sending async message", zap.Any("msg", msg))
	md, err := m.extractMetadata(ctx)
	if err != nil {
		return err
	}
	m.client.SendText(id.RoomID(md.RoomID), msg.Body) // TODO refactor
	return nil
}

func (m *messenger) Inbound() chan messaging.Message {
	return m.msgChannel
}

func (m *messenger) extractMetadata(ctx context.Context) (*messaging.Metadata, error) {
	mdPairs, ok := metadata.FromIncomingContext(ctx)
	metadata := &messaging.Metadata{}

	if !ok {
		return metadata, fmt.Errorf("metadata not found in context")
	}

	if sender, found := mdPairs["sender"]; found {
		metadata.Sender = sender[0]
	}

	if roomID, found := mdPairs["room_id"]; found {
		metadata.RoomID = roomID[0]
	}

	if cheques, found := mdPairs["cheques"]; found {
		chequesJSON := strings.Join(cheques, "")
		if err := json.Unmarshal([]byte(chequesJSON), &metadata.Cheques); err != nil {
			return metadata, fmt.Errorf("error unmarshalling cheques: %v", err)
		}
	}

	return metadata, nil
}