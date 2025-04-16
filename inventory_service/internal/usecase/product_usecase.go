package usecase

import (
	"context"
	"inventory_service/internal/domain"
)

type ProductUsecase struct {
	repo domain.ProductRepository
}

func NewProductUsecase(r domain.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: r}
}

func (uc *ProductUsecase) Create(ctx context.Context, p *domain.Product) error {
	return uc.repo.Create(ctx, p)
}

func (uc *ProductUsecase) GetByID(ctx context.Context, id int) (*domain.Product, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *ProductUsecase) Update(ctx context.Context, p *domain.Product) error {
	return uc.repo.Update(ctx, p)
}

func (uc *ProductUsecase) Delete(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *ProductUsecase) List(ctx context.Context) ([]domain.Product, error) {
	return uc.repo.List(ctx)
}
