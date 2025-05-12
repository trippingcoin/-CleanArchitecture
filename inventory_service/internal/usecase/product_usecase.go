package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"inventory_service/internal/cache"
	"inventory_service/internal/domain"

	"github.com/nats-io/nats.go"
)

type ProductUsecase struct {
	repo     domain.ProductRepository
	natsConn *nats.Conn
}

func NewProductUsecase(r domain.ProductRepository, nc *nats.Conn) *ProductUsecase {
	return &ProductUsecase{repo: r, natsConn: nc}
}

func (uc *ProductUsecase) publishEvent(subject string, data interface{}) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return uc.natsConn.Publish(subject, payload)
}

func (uc *ProductUsecase) Create(ctx context.Context, p *domain.Product) error {
	err := uc.repo.Create(ctx, p)
	if err != nil {
		return err
	}
	cache.InventoryCache.Set(fmt.Sprintf("product_%d", p.ID), p, cache.DefaultExpiration)
	uc.invalidateProductListCache()
	return uc.publishEvent("product.created", p)
}

func (uc *ProductUsecase) GetByID(ctx context.Context, id int) (*domain.Product, error) {
	cacheKey := fmt.Sprintf("product_%d", id)
	if cached, found := cache.InventoryCache.Get(cacheKey); found {
		return cached.(*domain.Product), nil
	}

	product, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	cache.InventoryCache.Set(cacheKey, product, cache.DefaultExpiration)
	return product, nil
}

func (uc *ProductUsecase) Update(ctx context.Context, p *domain.Product) error {
	err := uc.repo.Update(ctx, p)
	if err != nil {
		return err
	}
	cache.InventoryCache.Set(fmt.Sprintf("product_%d", p.ID), p, cache.DefaultExpiration)
	uc.invalidateProductListCache()
	return uc.publishEvent("product.updated", p)
}

func (uc *ProductUsecase) Delete(ctx context.Context, id int) error {
	err := uc.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	cache.InventoryCache.Delete(fmt.Sprintf("product_%d", id))
	uc.invalidateProductListCache()
	return uc.publishEvent("product.deleted", map[string]int{"id": id})
}

func (uc *ProductUsecase) List(ctx context.Context) ([]domain.Product, error) {
	const cacheKey = "product_list"
	if cached, found := cache.InventoryCache.Get(cacheKey); found {
		return cached.([]domain.Product), nil
	}

	products, err := uc.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	cache.InventoryCache.Set(cacheKey, products, cache.DefaultExpiration)
	return products, nil
}

func (uc *ProductUsecase) invalidateProductListCache() {
	cache.InventoryCache.Delete("product_list")
}
