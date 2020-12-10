package repository

import (
	"context"
	"ecommerce/app/models"
)

// ProductRepository ...
type ProductRepository interface {
	ResolveProductByID(ctx context.Context, id int) (*models.Product, error)
	ResolveProducts(ctx context.Context) ([]*models.Product, error)
	StoreProduct(ctx context.Context, product *models.Product) error
	UpdateProduct(ctx context.Context, id int, product *models.Product) error
	DeleteProductByID(ctx context.Context, id int) error
}
