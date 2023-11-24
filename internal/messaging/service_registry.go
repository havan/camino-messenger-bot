/*
 * Copyright (C) 2022-2023, Chain4Travel AG. All rights reserved.
 * See the file LICENSE for licensing terms.
 */

package messaging

import (
	"buf.build/gen/go/chain4travel/camino-messenger-protocol/grpc/go/cmp/services/accommodation/v1alpha1/accommodationv1alpha1grpc"
	"buf.build/gen/go/chain4travel/camino-messenger-protocol/grpc/go/cmp/services/ping/v1alpha1/pingv1alpha1grpc"
	"github.com/chain4travel/camino-messenger-bot/config"
	"github.com/chain4travel/camino-messenger-bot/internal/rpc/client"
	"go.uber.org/zap"
)

type ServiceRegistry struct {
	logger    *zap.SugaredLogger
	rpcClient *client.RPCClient
	services  map[MessageType]Service
}

func NewServiceRegistry(logger *zap.SugaredLogger, rpcClient *client.RPCClient) *ServiceRegistry {
	return &ServiceRegistry{
		logger:    logger,
		rpcClient: rpcClient,
		services:  make(map[MessageType]Service),
	}
}

func (s *ServiceRegistry) RegisterServices(requestTypes config.SupportedRequestTypesFlag) {
	for _, requestType := range requestTypes {
		var service Service
		switch MessageType(requestType) {
		case AccommodationSearchRequest:
			c := accommodationv1alpha1grpc.NewAccommodationSearchServiceClient(s.rpcClient.ClientConn)
			service = accommodationService{client: &c}
		case PingRequest:
			c := pingv1alpha1grpc.NewPingServiceClient(s.rpcClient.ClientConn)
			service = pingService{client: &c}
		default:
			s.logger.Infof("Skipping registration of unknown request type: %s", requestType)
			continue
		}
		s.services[MessageType(requestType)] = service
	}
}

func (s *ServiceRegistry) GetService(messageType MessageType) (Service, bool) {
	service, ok := s.services[messageType]
	return service, ok
}
