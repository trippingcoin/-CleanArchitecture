package domain

import "context"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	Stock       int
}

type ProductRepository interface {
	Create(ctx context.Context, p *Product) error
	GetByID(ctx context.Context, id int) (*Product, error)
	Update(ctx context.Context, p *Product) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]Product, error)
}
