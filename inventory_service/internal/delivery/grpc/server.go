package grpc

import (
	"context"
	"inventory_service/internal/domain"
	"inventory_service/internal/usecase"
	"inventory_service/proto/inventorypb"
)

type InventoryServer struct {
	inventorypb.UnimplementedInventoryServiceServer
	usecase *usecase.ProductUsecase
}

func NewInventoryServer(uc *usecase.ProductUsecase) *InventoryServer {
	return &InventoryServer{usecase: uc}
}

func (s *InventoryServer) CreateProduct(ctx context.Context, req *inventorypb.CreateProductRequest) (*inventorypb.Product, error) {
	product := &domain.Product{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       float32(req.GetPrice()),
		Stock:       int(req.GetStock()),
	}

	err := s.usecase.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return &inventorypb.Product{
		ProductId:   int32(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       int32(product.Stock),
	}, nil
}

func (s *InventoryServer) GetProduct(ctx context.Context, req *inventorypb.GetProductRequest) (*inventorypb.Product, error) {
	product, err := s.usecase.GetByID(ctx, int(req.GetProductId()))
	if err != nil {
		return nil, err
	}

	return &inventorypb.Product{
		ProductId:   int32(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       int32(product.Stock),
	}, nil
}

func (s *InventoryServer) ListProducts(ctx context.Context, _ *inventorypb.ListProductsRequest) (*inventorypb.ListProductsResponse, error) {
	products, err := s.usecase.List(ctx)
	if err != nil {
		return nil, err
	}

	var resp inventorypb.ListProductsResponse
	for _, p := range products {
		resp.Products = append(resp.Products, &inventorypb.Product{
			ProductId:   int32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       int32(p.Stock),
		})
	}

	return &resp, nil
}

func (s *InventoryServer) UpdateProduct(ctx context.Context, req *inventorypb.UpdateProductRequest) (*inventorypb.Product, error) {
	product := &domain.Product{
		ID:          int(req.GetProductId()),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       float32(req.GetPrice()),
		Stock:       int(req.GetStock()),
	}

	if err := s.usecase.Update(ctx, product); err != nil {
		return nil, err
	}

	return &inventorypb.Product{
		ProductId:   int32(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       int32(product.Stock),
	}, nil
}

func (s *InventoryServer) DeleteProduct(ctx context.Context, req *inventorypb.DeleteProductRequest) (*inventorypb.DeleteProductResponse, error) {
	if err := s.usecase.Delete(ctx, int(req.GetProductId())); err != nil {
		return nil, err
	}

	return &inventorypb.DeleteProductResponse{Success: true}, nil
}
