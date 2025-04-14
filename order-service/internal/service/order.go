package service

import (
	"order-service/internal/model"
	"order-service/internal/repository"
)

type OrderService interface {
	CreateOrder(order model.Order) (model.Order, error)
	GetOrderByID(id string) (model.Order, error)
	UpdateOrderStatus(id, status string) (model.Order, error)
	ListOrders(userID string) ([]model.Order, error)
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) CreateOrder(o model.Order) (model.Order, error) {
	return s.repo.Create(o)
}

func (s *orderService) GetOrderByID(id string) (model.Order, error) {
	return s.repo.GetByID(id)
}

func (s *orderService) UpdateOrderStatus(id, status string) (model.Order, error) {
	return s.repo.UpdateStatus(id, status)
}

func (s *orderService) ListOrders(userID string) ([]model.Order, error) {
	return s.repo.List(userID)
}
