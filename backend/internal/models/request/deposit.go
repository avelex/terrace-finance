package request

import "github.com/shopspring/decimal"

type DepositAndStake struct {
	Amount      decimal.Decimal `json:"amount"`
	Attestation string          `json:"attestation"`
	Signature   string          `json:"signature"`
}
