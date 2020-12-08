package product

import (
	"context"

	"gorm.io/gorm"
)

// Service ...
type Service struct {
	repo *Repository
}

// NewService ...
func NewService(db *gorm.DB) *Service {
	return &Service{
		repo: newRepository(db),
	}
}

// GetProductByID ...
func (s *Service) GetProductByID(ctx context.Context, id int) (Product, error) {
	return s.repo.GetProductByID(ctx, id)
}

// GetProducts ...
func (s *Service) GetProducts(ctx context.Context) ([]Product, error) {
	return s.repo.GetProducts(ctx)
}

// CreateProduct ...
func (s *Service) CreateProduct(ctx context.Context, product Product) (Product, error) {
	return s.repo.CreateProduct(ctx, product)
}

// UpdateProduct ...
func (s *Service) UpdateProduct(ctx context.Context, id int, product Product) (Product, error) {
	return s.repo.UpdateProduct(ctx, product)
}

// DeleteProduct ...
func (s *Service) DeleteProduct(ctx context.Context, id int) (Product, error) {
	return s.repo.DeleteProduct(ctx, id)
}
