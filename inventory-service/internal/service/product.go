package service

import (
	"inventory-service/internal/models"
	"inventory-service/internal/repository"
)

type InventoryService interface {
	CreateProduct(product models.Product) (models.Product, error)
	GetProductByID(id string) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(id string) error
	ListProducts() ([]models.Product, error)
}

type inventoryService struct {
	repo repository.ProductRepository
}

func NewInventoryService(repo repository.ProductRepository) InventoryService {
	return &inventoryService{repo: repo}
}

func (s *inventoryService) CreateProduct(p models.Product) (models.Product, error) {
	return s.repo.Create(p)
}

func (s *inventoryService) GetProductByID(id string) (models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *inventoryService) UpdateProduct(p models.Product) (models.Product, error) {
	return s.repo.Update(p)
}

func (s *inventoryService) DeleteProduct(id string) error {
	return s.repo.Delete(id)
}

func (s *inventoryService) ListProducts() ([]models.Product, error) {
	return s.repo.List()
}
