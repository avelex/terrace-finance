package models

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/uptrace/bun"
)

type AaveV3SupplySnapshot struct {
	bun.BaseModel `bun:"table:aave_v3_supply_snapshot"`

	ID        int64           `bun:",pk,autoincrement"`
	Network   string          `bun:"network,notnull"`
	APY       decimal.Decimal `bun:"apy,notnull"`
	Timestamp time.Time       `bun:"timestamp,notnull"`
}
