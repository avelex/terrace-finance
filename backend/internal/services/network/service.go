package network

import (
	"context"
	"fmt"
	"math/big"

	"github.com/avelex/terrace-finance/backend/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Service struct {
	clients map[string]*ethclient.Client
}

func NewService(ctx context.Context, cfg config.Config) (*Service, error) {
	clients := make(map[string]*ethclient.Client)
	for _, network := range cfg.Networks {
		client, err := ethclient.DialContext(ctx, network.URL)
		if err != nil {
			return nil, fmt.Errorf("dial context for %s: %w", network.Name, err)
		}
		clients[network.Name] = client
	}
	return &Service{
		clients: clients,
	}, nil
}

func (s *Service) Client(network string) (*ethclient.Client, error) {
	client, exists := s.clients[network]
	if !exists {
		return nil, fmt.Errorf("client for %s not found", network)
	}
	return client, nil
}

func (s *Service) ChainID(ctx context.Context, network string) (*big.Int, error) {
	client, err := s.Client(network)
	if err != nil {
		return nil, fmt.Errorf("client for %s: %w", network, err)
	}
	return client.ChainID(ctx)
}
