package strategy

import (
	"context"
	"fmt"

	"github.com/avelex/terrace-finance/backend/internal/models"

	"github.com/ethereum/go-ethereum/common"
)

var USDC_MAPPING = map[string]common.Address{
	"base":     common.HexToAddress("0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913"),
	"arbitrum": common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831"),
}

type Repository interface {
	InsertAaveV3SupplySnapshot(ctx context.Context, snapshot models.AaveV3SupplySnapshot) error
}

type Manager struct {
	repo Repository
}

func NewManager(repo Repository) *Manager {
	return &Manager{
		repo: repo,
	}
}

func (m *Manager) UpdateReserveData(ctx context.Context, reserve models.ReserveDataUpdated) error {
	// skip, if network is not supported
	usdc, exists := USDC_MAPPING[reserve.Network]
	if !exists {
		return nil
	}

	// skip, if reserve is not USDC
	if reserve.Reserve.Cmp(usdc) != 0 {
		return nil
	}

	if err := m.repo.InsertAaveV3SupplySnapshot(ctx, reserve.Snapshot()); err != nil {
		return fmt.Errorf("insert aave v3 supply snapshot: %w", err)
	}

	return nil
}
