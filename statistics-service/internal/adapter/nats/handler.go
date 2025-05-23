package nats

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/Neroframe/ecommerce-platform/statistics-service/internal/domain"
	"github.com/Neroframe/ecommerce-platform/statistics-service/internal/usecase"
	"github.com/nats-io/nats.go"
)

type StatisticsHandler struct {
	uc *usecase.StatisticsUsecase
	nc *nats.Conn
}

func NewStatisticsHandler(uc *usecase.StatisticsUsecase, nc *nats.Conn) *StatisticsHandler {
	return &StatisticsHandler{uc: uc, nc: nc}
}

func (h *StatisticsHandler) Handle(ctx context.Context, msg *nats.Msg) error {
	// decode
	var evt domain.Event
	if err := json.Unmarshal(msg.Data, &evt); err != nil {
		return fmt.Errorf("unmarshal event: %w", err)
	}
	if err := h.uc.HandleEvent(ctx, evt); err != nil {
		return fmt.Errorf("handle event: %w", err)
	}

	// acknowledgement
	var (
		ackSub string
		ackMsg string
	)
	if strings.HasPrefix(evt.EventType, "order.") {
		ackSub = "statistics.order.received"
		ackMsg = fmt.Sprintf("order with id %s is received", evt.EntityID)
	} else {
		ackSub = "statistics.inventory.received"
		ackMsg = fmt.Sprintf("item with id %s is received", evt.EntityID)
	}

	// publish acknowledgement
	if err := h.nc.Publish(ackSub, []byte(ackMsg)); err != nil {
		log.Printf("[NATS][ACK] failed to publish to %s: %v", ackSub, err)
	} else {
		log.Printf("[NATS][ACK] published to %s: %q", ackSub, ackMsg)
	}

	return nil
}
