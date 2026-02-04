package request

import "github.com/avelex/terrace-finance/backend/internal/models/enum"

type UnifyBalances struct {
	Domains []enum.CircleDomain `json:"domains"`
}
