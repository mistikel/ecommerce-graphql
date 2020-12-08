package product

import (
	"context"

	"gorm.io/gorm"
)

// Repository ...
type Repository struct {
	db *gorm.DB
}

// newRepository ...
func newRepository(db *gorm.DB) *Repository {
	r := &Repository{
		db: db,
	}
	return r
}

// GetProductByID ...
func (r *Repository) GetProductByID(ctx context.Context, id int) (Product, error) {
	var p Product
	result := r.db.Find(&p, id)
	if result.Error != nil {
		return Product{}, result.Error
	}
	return p, nil
}

// GetProducts ...
func (r *Repository) GetProducts(ctx context.Context) ([]Product, error) {
	var ps []Product
	result := r.db.Find(&ps)
	if result.Error != nil {
		return []Product{}, result.Error
	}
	return ps, nil
}

// CreateProduct ...
func (r *Repository) CreateProduct(ctx context.Context, product Product) (Product, error) {
	result := r.db.Create(&product)
	if result.Error != nil {
		return Product{}, result.Error
	}
	return product, nil
}

// UpdateProduct ...
func (r *Repository) UpdateProduct(ctx context.Context, product Product) (Product, error) {
	result := r.db.Save(&product)
	if result.Error != nil {
		return Product{}, result.Error
	}
	return product, nil
}

// DeleteProduct ...
func (r *Repository) DeleteProduct(ctx context.Context, id int) (Product, error) {
	result := r.db.Delete(&Product{}, id)
	if result.Error != nil {
		return Product{}, result.Error
	}
	return Product{}, nil
}
