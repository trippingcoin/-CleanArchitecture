package usecase

import (
	"context"
	"encoding/json"
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
	return uc.publishEvent("product.created", p)
}

func (uc *ProductUsecase) GetByID(ctx context.Context, id int) (*domain.Product, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *ProductUsecase) Update(ctx context.Context, p *domain.Product) error {
	err := uc.repo.Update(ctx, p)
	if err != nil {
		return err
	}
	return uc.publishEvent("product.updated", p)
}

func (uc *ProductUsecase) Delete(ctx context.Context, id int) error {
	err := uc.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return uc.publishEvent("product.deleted", map[string]int{"id": id})
}

func (uc *ProductUsecase) List(ctx context.Context) ([]domain.Product, error) {
	return uc.repo.List(ctx)
}
