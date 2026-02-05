package models

import (
	"time"

	"github.com/uptrace/bun"
)

type BridgeOp struct {
	bun.BaseModel `bun:"table:bridge_ops"`

	ID             string `bun:"id,pk" json:"id"`
	FromTerrace    string `bun:"from_terrace,notnull" json:"fromTerrace"`
	FromDomain     uint32 `bun:"from_domain,notnull" json:"fromDomain"`
	ToTerrace      string `bun:"to_terrace,notnull" json:"toTerrace"`
	ToDomain       uint32 `bun:"to_domain,notnull" json:"toDomain"`
	SendAmount     string `bun:"send_amount,notnull" json:"sendAmount"`
	ReceivedAmount string `bun:"received_amount,nullzero" json:"receivedAmount"`
	SentTxHash     string `bun:"sent_tx_hash,notnull" json:"sentTxHash"`
	ReceivedTxHash string `bun:"received_tx_hash,nullzero" json:"receivedTxHash"`

	CCTPMessage     string `bun:"cctp_message,nullzero" json:"cctpMessage"`
	CCTPAttestation string `bun:"cctp_attestation,nullzero" json:"cctpAttestation"`

	SentAt     time.Time `bun:"sent_at,notnull" json:"sentAt"`
	ReceivedAt time.Time `bun:"received_at,nullzero" json:"receivedAt"`
}
