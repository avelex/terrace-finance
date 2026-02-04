package response

import (
	"github.com/avelex/terrace-finance/backend/internal/models/enum"

	"github.com/shopspring/decimal"
)

type ProtocolBalances struct {
	StakingBalance    decimal.Decimal                       `json:"stakingBalance"`
	StablecoinBalance decimal.Decimal                       `json:"stablecoinBalance"`
	UnifiedUSDC       map[enum.CircleDomain]decimal.Decimal `json:"unifiedUsdc"`
	USDC              map[enum.CircleDomain]decimal.Decimal `json:"usdc"`
}
