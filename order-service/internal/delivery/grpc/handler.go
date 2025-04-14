package grpcH

import (
	"context"
	"order-service/internal/model"
	"order-service/internal/service"
	pb "order-service/order-service/proto"
)

type OrderHandler struct {
	pb.UnimplementedOrderServiceServer
	svc service.OrderService
}

func NewOrderHandler(svc service.OrderService) *OrderHandler {
	return &OrderHandler{svc: svc}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	order := model.Order{
		UserID: req.GetUserId(),
		Items:  convertToModelItems(req.GetItems()),
	}

	created, err := h.svc.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	return convertToProtoOrder(created), nil
}

func (h *OrderHandler) GetOrderByID(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	order, err := h.svc.GetOrderByID(req.GetId())
	if err != nil {
		return nil, err
	}
	return convertToProtoOrder(order), nil
}

func (h *OrderHandler) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest) (*pb.OrderResponse, error) {
	order, err := h.svc.UpdateOrderStatus(req.GetId(), req.GetStatus())
	if err != nil {
		return nil, err
	}
	return convertToProtoOrder(order), nil
}

func (h *OrderHandler) ListUserOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := h.svc.ListOrders(req.GetUserId())
	if err != nil {
		return nil, err
	}

	var res pb.ListOrdersResponse
	for _, order := range orders {
		res.Orders = append(res.Orders, convertToProtoOrder(order))
	}
	return &res, nil
}

// Helpers
func convertToModelItems(items []*pb.OrderItem) []model.OrderItem {
	var result []model.OrderItem
	for _, i := range items {
		result = append(result, model.OrderItem{
			ProductID: i.GetProductId(),
			Quantity:  i.GetQuantity(),
		})
	}
	return result
}

func convertToProtoOrder(o model.Order) *pb.OrderResponse {
	var items []*pb.OrderItem
	for _, i := range o.Items {
		items = append(items, &pb.OrderItem{
			ProductId: i.ProductID,
			Quantity:  i.Quantity,
		})
	}
	return &pb.OrderResponse{
		Id:     o.ID,
		UserId: o.UserID,
		Status: o.Status,
		Items:  items,
	}
}
