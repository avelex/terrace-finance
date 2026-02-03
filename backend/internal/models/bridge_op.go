package models

import (
	"time"

	"github.com/uptrace/bun"
)

type BridgeOp struct {
	bun.BaseModel `bun:"table:bridge_ops"`

	ID             string `bun:"id,pk"`
	FromTerrace    string `bun:"from_terrace,notnull"`
	FromDomain     uint32 `bun:"from_domain,notnull"`
	ToTerrace      string `bun:"to_terrace,notnull"`
	ToDomain       uint32 `bun:"to_domain,notnull"`
	SendAmount     string `bun:"send_amount,notnull"`
	ReceivedAmount string `bun:"received_amount,nullzero"`
	SentTxHash     string `bun:"sent_tx_hash,notnull"`
	ReceivedTxHash string `bun:"received_tx_hash,nullzero"`

	CCTPMessage     string `bun:"cctp_message,nullzero"`
	CCTPAttestation string `bun:"cctp_attestation,nullzero"`

	SentAt     time.Time `bun:"sent_at,notnull"`
	ReceivedAt time.Time `bun:"received_at,nullzero"`
}
