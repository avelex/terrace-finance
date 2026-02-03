package repository

import (
	"context"
	"fmt"

	"github.com/avelex/terrace-finance/backend/internal/models"

	"github.com/uptrace/bun"
)

type AaveRepository struct {
	db *bun.DB
}

func NewAaveRepository(db *bun.DB) *AaveRepository {
	return &AaveRepository{
		db: db,
	}
}

func (r *AaveRepository) InsertAaveV3SupplySnapshot(ctx context.Context, snapshot models.AaveV3SupplySnapshot) error {
	_, err := r.db.NewInsert().Model(&snapshot).On("CONFLICT (id) DO NOTHING").Exec(ctx)
	if err != nil {
		return fmt.Errorf("unable to insert aave v3 supply snapshot: %w", err)
	}
	return nil
}
