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
	// Create an order in the repository.
	Create(order *Order) error

	// GetByID retrieves an order by its ID.
	GetByID(id string) (*Order, error)

	// UpdateStatus updates the status of an order.
	UpdateStatus(id string, status string) error

	// ListByUser retrieves a list of orders for a specific user.
	ListByUser(userID string) ([]Order, error)
}
