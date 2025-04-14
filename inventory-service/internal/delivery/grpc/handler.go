package grpcH

import (
	"context"
	"inventory-service/internal/models"
	"inventory-service/internal/service"

	pb "inventory-service/inventory-service/proto"
)

type InventoryHandler struct {
	pb.UnimplementedInventoryServiceServer
	svc service.InventoryService
}

func NewInventoryHandler(svc service.InventoryService) *InventoryHandler {
	return &InventoryHandler{svc: svc}
}

func (h *InventoryHandler) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	product := models.Product{
		Name:     req.GetName(),
		Category: req.GetCategory(),
		Price:    req.GetPrice(),
		Stock:    req.GetStock(),
	}
	created, _ := h.svc.CreateProduct(product)

	return &pb.ProductResponse{
		Id:       created.ID,
		Name:     created.Name,
		Category: created.Category,
		Price:    created.Price,
		Stock:    created.Stock,
	}, nil
}
func (h *InventoryHandler) GetProductByID(ctx context.Context, req *pb.GetProductRequest) (*pb.ProductResponse, error) {
	p, err := h.svc.GetProductByID(req.GetId())
	if err != nil {
		return nil, err
	}
	return convertToProto(p), nil
}

func (h *InventoryHandler) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.ProductResponse, error) {
	p := models.Product{
		ID:       req.GetId(),
		Name:     req.GetName(),
		Category: req.GetCategory(),
		Price:    req.GetPrice(),
		Stock:    req.GetStock(),
	}
	updated, err := h.svc.UpdateProduct(p)
	if err != nil {
		return nil, err
	}
	return convertToProto(updated), nil
}

func (h *InventoryHandler) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.Empty, error) {
	err := h.svc.DeleteProduct(req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (h *InventoryHandler) ListProducts(ctx context.Context, _ *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	products, _ := h.svc.ListProducts()
	res := &pb.ListProductsResponse{}
	for _, p := range products {
		res.Products = append(res.Products, convertToProto(p))
	}
	return res, nil
}

func convertToProto(p models.Product) *pb.ProductResponse {
	return &pb.ProductResponse{
		Id:       p.ID,
		Name:     p.Name,
		Category: p.Category,
		Price:    p.Price,
		Stock:    p.Stock,
	}
}
