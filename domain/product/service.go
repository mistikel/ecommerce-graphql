package product

import "gorm.io/gorm"

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
func (s *Service) GetProductByID(id int) (Product, error) {
	return s.repo.GetProductByID(id)
}

// GetProducts ...
func (s *Service) GetProducts() ([]Product, error) {
	return s.repo.GetProducts()
}

// CreateProduct ...
func (s *Service) CreateProduct(product Product) (Product, error) {
	return s.repo.CreateProduct(product)
}

// UpdateProduct ...
func (s *Service) UpdateProduct(product Product) (Product, error) {
	return s.repo.UpdateProduct(product)
}

// DeleteProduct ...
func (s *Service) DeleteProduct(id int) (Product, error) {
	return s.repo.DeleteProduct(id)
}
