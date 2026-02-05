package response

import (
	"github.com/avelex/terrace-finance/backend/internal/models/enum"
	"github.com/ethereum/go-ethereum/common"

	"github.com/shopspring/decimal"
)

type ProtocolBalances struct {
	StakingBalance    decimal.Decimal                       `json:"stakingBalance"`
	StablecoinBalance decimal.Decimal                       `json:"stablecoinBalance"`
	UnifiedUSDC       map[enum.CircleDomain]decimal.Decimal `json:"unifiedUsdc"`
	USDC              map[enum.CircleDomain]decimal.Decimal `json:"usdc"`
}

type VaultInfo struct {
	Domain  enum.CircleDomain `json:"domain"`
	Address common.Address    `json:"address"`
	Balance decimal.Decimal   `json:"balance"`
}
