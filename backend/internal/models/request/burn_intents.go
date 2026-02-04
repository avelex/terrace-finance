package request

import "github.com/avelex/terrace-finance/backend/internal/models"

type BurnIntents struct {
	Intents []models.BurnIntent `json:"burnIntents"`
}
