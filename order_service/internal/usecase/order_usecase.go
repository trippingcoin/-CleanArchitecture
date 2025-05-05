package usecase

import (
	"context"
	"encoding/json"
	"order_service/internal/domain"

	"github.com/nats-io/nats.go"
)

type OrderUsecase struct {
	orderRepo domain.OrderRepository
	natsConn  *nats.Conn
}

func NewOrderUsecase(orderRepo domain.OrderRepository, natsConn *nats.Conn) *OrderUsecase {
	return &OrderUsecase{
		orderRepo: orderRepo,
		natsConn:  natsConn,
	}
}

func (u *OrderUsecase) CreateOrder(ctx context.Context, userID string, items []domain.OrderItem, totalPrice float64) (string, error) {
	// Create an Order struct with all necessary fields
	order := &domain.Order{
		UserID:     userID,
		Items:      items,
		TotalPrice: totalPrice,
		Status:     "pending", // Setting the initial status to "pending"
	}

	// Pass the order to the repository layer for saving it to the database
	err := u.orderRepo.Create(order)
	if err != nil {
		return "", err
	}

	event := map[string]interface{}{
		"event":       "OrderCreated",
		"order_id":    order.ID,
		"user_id":     order.UserID,
		"total_price": order.TotalPrice,
		"created_at":  order.CreatedAt,
	}

	data, err := json.Marshal(event)
	if err != nil {
		return "error: ", err
	}
	err = u.natsConn.Publish("order.created", data)
	if err != nil {
		return "error: ", err
	}

	return order.ID, nil
}

func (u *OrderUsecase) GetOrder(ctx context.Context, orderID string) (*domain.Order, error) {
	return u.orderRepo.GetByID(orderID)
}

func (u *OrderUsecase) ListOrders(ctx context.Context, userID string) ([]domain.Order, error) {
	return u.orderRepo.ListByUser(userID)
}

func (u *OrderUsecase) UpdateOrderStatus(ctx context.Context, orderID string, status string) error {
	// Update the order's status
	return u.orderRepo.UpdateStatus(orderID, status)
}
