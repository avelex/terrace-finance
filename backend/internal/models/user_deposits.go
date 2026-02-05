package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserDeposit struct {
	bun.BaseModel `bun:"table:user_deposits"`

	ID              uuid.UUID `bun:"id,pk" json:"id"`
	Address         string    `bun:"address,notnull" json:"address"`
	Value           string    `bun:"value,notnull" json:"value"`
	DestDomain      uint32    `bun:"dest_domain,notnull" json:"destDomain"`
	DestGatewayMint string    `bun:"dest_gateway_mint,notnull" json:"destGatewayMint"`
	Attestation     string    `bun:"attestation,notnull" json:"attestation"`
	Signature       string    `bun:"signature,notnull" json:"signature"`
	TxHash          string    `bun:"tx_hash,nullzero" json:"txHash"`
	Reason          string    `bun:"reason,nullzero" json:"reason"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:current_timestamp" json:"createdAt"`
	DepositedAt     time.Time `bun:"deposited_at,nullzero" json:"depositedAt"`
}

type UserUnifiedPermits struct {
	bun.BaseModel `bun:"table:user_unified_permits"`

	ID            uuid.UUID `bun:"id" json:"id"`
	Owner         string    `bun:"owner,notnull" json:"owner"`
	Token         string    `bun:"token,notnull" json:"token"`
	Value         string    `bun:"value,notnull" json:"value"`
	Deadline      int64     `bun:"deadline,notnull" json:"deadline"`
	Signature     string    `bun:"signature,nullzero" json:"signature"`
	Domain        uint32    `bun:"domain,notnull" json:"domain"`
	GatewayWallet string    `bun:"gateway_wallet,notnull" json:"gatewayWallet"`
	TxHash        string    `bun:"tx_hash,nullzero" json:"txHash"`
	CreatedAt     time.Time `bun:"created_at,notnull,default:current_timestamp" json:"createdAt"`
	DepositedAt   time.Time `bun:"deposited_at,nullzero" json:"depositedAt"`
}
