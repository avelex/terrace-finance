package event_handler

import (
	"context"

	"github.com/avelex/terrace-finance/backend/internal/models"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"

	eventscale "github.com/eventscale/eventscale/pkg/sdk-go"
)

type Repository interface {
	CreateBridgeOp(ctx context.Context, op models.SendFunds) error
	CompleteBridgeOp(ctx context.Context, op models.ReceivedFunds) error
}

type StrategyManager interface {
	UpdateReserveData(ctx context.Context, reserve models.ReserveDataUpdated) error
}

type EventHandler struct {
	repo    Repository
	manager StrategyManager
}

func NewEventHandler(repo Repository, manager StrategyManager) *EventHandler {
	return &EventHandler{
		repo:    repo,
		manager: manager,
	}
}

func (h *EventHandler) Handle(ctx context.Context, block eventscale.EventBlock) error {
	for _, v := range block.Events {
		location := models.NewEventLocation(
			v.MetaData.TxHash,
			v.MetaData.Timestamp,
			v.MetaData.Network,
		)

		switch v.MetaData.Name {
		case "SendFunds":
			sendFunds := models.SendFunds{
				EventLocation: location,
			}

			if err := v.Decode(&sendFunds); err != nil {
				log.Error().Err(err).Msg("decode SendFunds event")
				return err
			}

			if err := h.handleSendFunds(ctx, sendFunds); err != nil {
				log.Error().Err(err).Msg("handle SendFunds event")
				return err
			}
		case "ReceivedFunds":
			receivedFunds := models.ReceivedFunds{
				EventLocation: location,
			}

			if err := v.Decode(&receivedFunds); err != nil {
				log.Error().Err(err).Msg("decode ReceivedFunds event")
				return err
			}

			if err := h.handleReceivedFunds(ctx, receivedFunds); err != nil {
				log.Error().Err(err).Msg("handle ReceivedFunds event")
				return err
			}
		case "ReserveDataUpdated":
			reserveDataUpdated := models.ReserveDataUpdated{
				EventLocation: location,
			}

			if err := v.Decode(&reserveDataUpdated); err != nil {
				log.Error().Err(err).Msg("decode ReserveDataUpdated event")
				return err
			}

			if err := h.manager.UpdateReserveData(ctx, reserveDataUpdated); err != nil {
				log.Error().Err(err).Msg("handle ReserveDataUpdated event")
				return err
			}
		}
	}

	return nil
}

func (h *EventHandler) handleSendFunds(ctx context.Context, event models.SendFunds) error {
	log.Info().
		Str("event_id", common.BytesToHash(event.ID[:]).String()).
		Str("tx_hash", event.TxHash.String()).
		Uint32("domain", event.FromDomain).
		Msgf("received SendFunds")

	return h.repo.CreateBridgeOp(ctx, event)
}

func (h *EventHandler) handleReceivedFunds(ctx context.Context, event models.ReceivedFunds) error {
	log.Info().
		Str("event_id", common.BytesToHash(event.ID[:]).String()).
		Str("tx_hash", event.TxHash.String()).
		Uint32("domain", event.ToDomain).
		Msgf("received ReceivedFunds")

	return h.repo.CompleteBridgeOp(ctx, event)
}
