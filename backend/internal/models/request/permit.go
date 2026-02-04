package request

import (
	"github.com/avelex/terrace-finance/backend/internal/models/enum"

	"github.com/google/uuid"
)

type SignedPermits struct {
	ID      uuid.UUID                    `json:"id"`
	Domains map[enum.CircleDomain]string `json:"domains"`
}
