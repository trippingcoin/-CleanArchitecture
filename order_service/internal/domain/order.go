package domain

import "time"

type Order struct {
	ID         string      `json:"order_id"`
	UserID     string      `json:"user_id"`
	TotalPrice float64     `json:"total_price"`
	Status     string      `json:"status"`
	Items      []OrderItem `json:"items"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type OrderItem struct {
	ProductID    string
	Quantity     int
	PricePerItem float64
}

type OrderRepository interface {
	Create(order *Order) error

	GetByID(id string) (*Order, error)

	UpdateStatus(id string, status string) error

	ListByUser(userID string) ([]Order, error)

	GetAll() ([]Order, error)

	GetOrderItems(orders string) ([]OrderItem, error)
}
