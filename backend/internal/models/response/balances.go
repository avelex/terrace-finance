package response

import "github.com/shopspring/decimal"

type ProtocolBalances struct {
	StakingBalance    decimal.Decimal            `json:"stakingBalance"`
	StablecoinBalance decimal.Decimal            `json:"stablecoinBalance"`
	UnifiedUSDC       map[uint32]decimal.Decimal `json:"unifiedUsdc"`
	USDC              map[uint32]decimal.Decimal `json:"usdc"`
}
