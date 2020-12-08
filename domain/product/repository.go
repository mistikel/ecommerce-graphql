package product

import (
	"ecommerce/database"

	"gorm.io/gorm"
)

// Repository ...
type Repository struct {
	db *gorm.DB
}

// newRepository ...
func newRepository() *Repository {
	r := &Repository{
		db: database.Get(),
	}
	r.db.AutoMigrate(&Variant{}, &Shipping{}, &Seo{}, &Product{})
	return r
}

// GetProductByID ...
func (r *Repository) GetProductByID(id int) (Product, error) {
	var p Product
	result := r.db.Find(&p, id)
	if result.Error != nil {
		return Product{}, result.Error
	}
	return p, nil
}

// GetProducts ...
func (r *Repository) GetProducts() ([]Product, error) {
	var ps []Product
	result := r.db.Find(&ps)
	if result.Error != nil {
		return []Product{}, result.Error
	}
	return ps, nil
}

// CreateProduct ...
func (r *Repository) CreateProduct(product Product) (Product, error) {
	result := r.db.Create(&product)
	if result.Error != nil {
		return Product{}, result.Error
	}
	return product, nil
}

// UpdateProduct ...
func (r *Repository) UpdateProduct(product Product) (Product, error) {
	result := r.db.Save(&product)
	if result.Error != nil {
		return Product{}, result.Error
	}
	return product, nil
}

// DeleteProduct ...
func (r *Repository) DeleteProduct(id int) (Product, error) {
	result := r.db.Delete(&Product{}, id)
	if result.Error != nil {
		return Product{}, result.Error
	}
	return Product{}, nil
}
