package nats

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Neroframe/ecommerce-platform/statistics-service/internal/usecase"
	"github.com/nats-io/nats.go"
)

type StatisticsHandler struct {
	uc usecase.StatisticsUsecase
}

func NewStatisticsHandler(uc usecase.StatisticsUsecase) *StatisticsHandler {
	return &StatisticsHandler{uc: uc}
}

// HandleOrderCreated only deals with the "order.created" subject.
func (h *StatisticsHandler) HandleOrderCreated(ctx context.Context, msg *nats.Msg) error {
	log.Printf("[NATS] Received subject=%s data=%s", msg.Subject, string(msg.Data))
	var evt usecase.OrderCreatedEvent
	if err := json.Unmarshal(msg.Data, &evt); err != nil {
		log.Printf("unmarshal order.created: %v", err)
		return err
	}

	err := h.uc.HandleOrderCreated(ctx, evt)
	if err != nil {
		log.Printf("[NATS] HandleOrderCreated error: %v", err)
		return err
	}

	log.Printf("[NATS] HandleOrderCreated succeeded for order_id=%s", evt.OrderID)
	return nil
}

// HandleUserRegistered only deals with the "user.registered" subject.
func (h *StatisticsHandler) HandleUserRegistered(ctx context.Context, msg *nats.Msg) error {
	var evt usecase.UserRegisteredEvent
	if err := json.Unmarshal(msg.Data, &evt); err != nil {
		log.Printf("unmarshal user.registered: %v", err)
		return err
	}
	return h.uc.HandleUserRegistered(ctx, evt)
}
