package domain

type Order struct {
	ID         string
	UserID     string
	Items      []OrderItem
	TotalPrice float64
	Status     string
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
