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
	UserID     string    `json:"user_id"`
	ProductID  string    `json:"id"`
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
	CategoryID string    `json:"category_id"`
	Timestamp  time.Time `json:"timestamp"`
}

type ProductUpdatedEvent struct {
	UserID     string    `json:"user_id"`
	ProductID  string    `json:"id"`
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
	CategoryID string    `json:"category_id"`
	Timestamp  time.Time `json:"timestamp"`
}

type ProductDeletedEvent struct {
	UserID    string    `json:"user_id"`
	ProductID string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}


// logging
func (e OrderCreatedEvent) GetOrderID() string  { return e.OrderID }
func (e OrderUpdatedEvent) GetOrderID() string  { return e.OrderID }
func (e OrderDeletedEvent) GetOrderID() string  { return e.OrderID }

func (e ProductCreatedEvent) GetProductID() string  { return e.ProductID }
func (e ProductUpdatedEvent) GetProductID() string  { return e.ProductID }
func (e ProductDeletedEvent) GetProductID() string  { return e.ProductID }
