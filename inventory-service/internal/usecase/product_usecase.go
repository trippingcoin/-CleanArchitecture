package usecase

import (
	"context"

	"inventory-service/internal/models"
)

type ProductRepo interface {
	Create(ctx context.Context, p *models.Product) (*models.Product, error)
	GetByID(ctx context.Context, id string) (*models.Product, error)
	Update(ctx context.Context, p *models.Product) (*models.Product, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*models.Product, error)
}

type ProductUsecase struct {
	Repo ProductRepo
}

func NewProductUsecase(r ProductRepo) *ProductUsecase {
	return &ProductUsecase{Repo: r}
}

func (u *ProductUsecase) CreateProduct(ctx context.Context, p *models.Product) (*models.Product, error) {
	return u.Repo.Create(ctx, p)
}

func (u *ProductUsecase) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	return u.Repo.GetByID(ctx, id)
}

func (u *ProductUsecase) UpdateProduct(ctx context.Context, p *models.Product) (*models.Product, error) {
	return u.Repo.Update(ctx, p)
}

func (u *ProductUsecase) DeleteProduct(ctx context.Context, id string) error {
	return u.Repo.Delete(ctx, id)
}

func (u *ProductUsecase) ListProducts(ctx context.Context) ([]*models.Product, error) {
	return u.Repo.List(ctx)
}
