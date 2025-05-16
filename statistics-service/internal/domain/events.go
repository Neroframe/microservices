package domain

import "time"

type OrderCreatedEvent struct {
	UserID    string    `json:"UserID"`
	OrderID   string    `json:"OrderID"`
	Timestamp time.Time `json:"Timestamp"`
}
type OrderUpdatedEvent struct {
	UserID    string    `json:"UserID"`
	OrderID   string    `json:"OrderID"`
	Timestamp time.Time `json:"Timestamp"`
}

type OrderDeletedEvent struct {
	UserID    string    `json:"UserID"`
	OrderID   string    `json:"OrderID"`
	Timestamp time.Time `json:"Timestamp"`
}

type ProductCreatedEvent struct {
	UserID    string    `json:"UserID"`
	ProductID string    `json:"ProductID"`
	Timestamp time.Time `json:"Timestamp"`
}

type ProductUpdatedEvent struct {
	UserID    string    `json:"UserID"`
	ProductID string    `json:"ProductID"`
	Timestamp time.Time `json:"Timestamp"`
}

type ProductDeletedEvent struct {
	UserID    string    `json:"UserID"`
	ProductID string    `json:"ProductID"`
	Timestamp time.Time `json:"Timestamp"`
}

// logging
func (e OrderCreatedEvent) GetOrderID() string  { return e.OrderID }
func (e OrderUpdatedEvent) GetOrderID() string  { return e.OrderID }
func (e OrderDeletedEvent) GetOrderID() string  { return e.OrderID }

func (e ProductCreatedEvent) GetProductID() string  { return e.ProductID }
func (e ProductUpdatedEvent) GetProductID() string  { return e.ProductID }
func (e ProductDeletedEvent) GetProductID() string  { return e.ProductID }
