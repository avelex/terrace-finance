package request

import "github.com/google/uuid"

type DepositAndStake struct {
	ID          uuid.UUID `json:"id"`
	Attestation string    `json:"attestation"`
	Signature   string    `json:"signature"`
}
