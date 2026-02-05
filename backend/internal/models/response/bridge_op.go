package response

import "github.com/avelex/terrace-finance/backend/internal/models"

type PagableBridgeOps struct {
	Data  []models.BridgeOp `json:"data"`
	Total int64             `json:"total"`
	Page  int64             `json:"page"`
	Limit int64             `json:"limit"`
}

func NewPagableBridgeOps(ops []models.BridgeOp, total int64, page, limit int64) PagableBridgeOps {
	return PagableBridgeOps{
		Data:  ops,
		Total: total,
		Page:  page,
		Limit: limit,
	}
}
