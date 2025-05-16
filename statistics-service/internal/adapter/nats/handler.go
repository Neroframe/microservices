package nats

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Neroframe/ecommerce-platform/statistics-service/internal/domain"
	"github.com/Neroframe/ecommerce-platform/statistics-service/internal/usecase"
	"github.com/nats-io/nats.go"
)

type StatisticsHandler struct {
	uc *usecase.StatisticsUsecase
}

func NewStatisticsHandler(uc *usecase.StatisticsUsecase) *StatisticsHandler {
	return &StatisticsHandler{uc: uc}
}

func (h *StatisticsHandler) HandleOrderCreated(ctx context.Context, msg *nats.Msg) error {
	var evt domain.OrderCreatedEvent
	return handleEvent(ctx, msg, "HandleOrderCreated", evt, h.uc.HandleOrderCreated)
}

func (h *StatisticsHandler) HandleOrderUpdated(ctx context.Context, msg *nats.Msg) error {
	var evt domain.OrderUpdatedEvent
	return handleEvent(ctx, msg, "HandleOrderUpdated", evt, h.uc.HandleOrderUpdated)
}

func (h *StatisticsHandler) HandleOrderDeleted(ctx context.Context, msg *nats.Msg) error {
	var evt domain.OrderDeletedEvent
	return handleEvent(ctx, msg, "HandleOrderDeleted", evt, h.uc.HandleOrderDeleted)
}

func (h *StatisticsHandler) HandleProductCreated(ctx context.Context, msg *nats.Msg) error {
	var evt domain.ProductCreatedEvent
	return handleEvent(ctx, msg, "HandleProductCreated", evt, h.uc.HandleProductCreated)
}

func (h *StatisticsHandler) HandleProductUpdated(ctx context.Context, msg *nats.Msg) error {
	var evt domain.ProductUpdatedEvent
	return handleEvent(ctx, msg, "HandleProductUpdated", evt, h.uc.HandleProductUpdated)
}

func (h *StatisticsHandler) HandleProductDeleted(ctx context.Context, msg *nats.Msg) error {
	var evt domain.ProductDeletedEvent
	return handleEvent(ctx, msg, "HandleProductDeleted", evt, h.uc.HandleProductDeleted)
}

func handleEvent[T any](ctx context.Context, msg *nats.Msg, logPrefix string, unmarshalTo T, handler func(context.Context, T) error) error {
	log.Printf("[NATS] Received subject=%s data=%s", msg.Subject, string(msg.Data))

	if err := json.Unmarshal(msg.Data, &unmarshalTo); err != nil {
		log.Printf("unmarshal %s: %v", logPrefix, err)
		return err
	}

	if err := handler(ctx, unmarshalTo); err != nil {
		log.Printf("[NATS] %s error: %v", logPrefix, err)
		return err
	}

	// Proper interface-based logging
	switch v := any(unmarshalTo).(type) {
	case interface{ GetOrderID() string }:
		log.Printf("[NATS] %s succeeded for order_id=%s", logPrefix, v.GetOrderID())
	case interface{ GetProductID() string }:
		log.Printf("[NATS] %s succeeded for product_id=%s", logPrefix, v.GetProductID())
	default:
		log.Printf("[NATS] %s succeeded", logPrefix)
	}

	return nil
}
