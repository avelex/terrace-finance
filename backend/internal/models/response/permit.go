package response

import "github.com/avelex/terrace-finance/backend/internal/models/enum"

type PermitPayload struct {
	ID      string                       `json:"id"`
	Domains map[enum.CircleDomain]string `json:"domains"`
}
