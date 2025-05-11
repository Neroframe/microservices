package nats

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Neroframe/ecommerce-platform/statistics-service/internal/usecase"
	"github.com/nats-io/nats.go"
)

// Handler routes different subjects to the right use-case methods.
type Handler struct {
	uc usecase.StatisticsUsecase
}

// NewStatisticsHandler binds your usecase.
func NewStatisticsHandler(uc usecase.StatisticsUsecase) *Handler {
	return &Handler{uc: uc}
}

// Handle gets invoked for every subscribed subject.
func (h *Handler) Handle(ctx context.Context, msg *nats.Msg) error {
	log.Printf("[NATS] %s → %s", msg.Subject, string(msg.Data))

	switch msg.Subject {
	case "order.created":
		var evt OrderCreatedEvent
		if err := json.Unmarshal(msg.Data, &evt); err != nil {
			return err
		}
		// delegate to usecase (you’d add this method)
		return h.uc.HandleOrderCreated(ctx, evt.UserID, evt.OrderID, evt.Timestamp)

	// case "order.deleted":
	//     // similar

	// case "inventory.updated":
	//     // parse inventory event and call h.uc.HandleInventoryUpdated...

	default:
		log.Printf("unhandled subject: %s", msg.Subject)
		return nil
	}
}
