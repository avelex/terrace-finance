package repository

import (
	"context"
	"fmt"

	"github.com/avelex/terrace-finance/backend/internal/models"

	"github.com/uptrace/bun"
)

type StrategyRepository struct {
	db *bun.DB
}

func NewStrategyRepository(db *bun.DB) *StrategyRepository {
	return &StrategyRepository{db: db}
}

func (r *StrategyRepository) GetByDomainAndName(ctx context.Context, domain uint32, name string) (*models.Strategy, error) {
	var strategy models.Strategy
	err := r.db.NewSelect().Model(&strategy).
		Where("domain = ?", domain).
		Where("name = ?", name).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("strategy not found: %w", err)
	}

	return &strategy, nil
}
