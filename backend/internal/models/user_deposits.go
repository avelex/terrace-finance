package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserDeposit struct {
	bun.BaseModel `bun:"table:user_deposits"`

	ID      uuid.UUID `bun:"id,pk"`
	Address string    `bun:"address,notnull"`
	Value   string    `bun:"value,notnull"`

	DestDomain      uint32 `bun:"dest_domain,notnull"`
	DestGatewayMint string `bun:"dest_gateway_mint,notnull"`
	Attestation     string `bun:"attestation,nullzero"`
	Signature       string `bun:"signature,nullzero"`
	DepositTxHash   string `bun:"deposit_tx_hash,nullzero"`

	CreatedAt   time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UnifiedAt   time.Time `bun:"unified_at,nullzero"`
	DepositedAt time.Time `bun:"deposited_at,nullzero"`
}

type UserUnifiedPermits struct {
	bun.BaseModel `bun:"table:user_unified_permits"`

	DepositID     uuid.UUID `bun:"deposit_id,notnull"`
	Owner         string    `bun:"owner,notnull"`
	Token         string    `bun:"token,notnull"`
	Value         string    `bun:"value,notnull"`
	Deadline      int64     `bun:"deadline,notnull"`
	Signature     string    `bun:"signature,nullzero"`
	Domain        uint32    `bun:"domain,notnull"`
	GatewayWallet string    `bun:"gateway_wallet,notnull"`
	TxHash        string    `bun:"tx_hash,nullzero"`
	CreatedAt     time.Time `bun:"created_at,notnull,default:current_timestamp"`
	DepositedAt   time.Time `bun:"deposited_at,nullzero"`
}
