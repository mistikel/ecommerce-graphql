package service

import (
	"context"
	"ecommerce/app/domain/repository"
	repo "ecommerce/app/infrastructure/repository"
	"ecommerce/app/models"
	"log"

	"gorm.io/gorm"
)

// ProductService ...
type ProductService struct {
	repo repository.ProductRepository
}

// NewProductService ...
func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{
		repo: repo.NewProductRepository(db),
	}
}

// CreateProduct ...
func (p *ProductService) CreateProduct(ctx context.Context, input models.Input) (*models.Product, error) {
	product := input.ToProduct()
	err := p.repo.StoreProduct(ctx, &product)
	if err != nil {
		return nil, err
	}
	return p.repo.ResolveProductByID(ctx, product.ID)
}

// UpdateProduct ...
func (p *ProductService) UpdateProduct(ctx context.Context, id int, input models.Input) (*models.Product, error) {
	oldProduct, err := p.repo.ResolveProductByID(ctx, id)
	if err != nil || oldProduct == nil {
		log.Print("product not exist")
		return nil, err
	}
	product := input.ToProduct()
	err = p.repo.UpdateProduct(ctx, id, &product)
	if err != nil {
		return nil, err
	}
	return p.repo.ResolveProductByID(ctx, product.ID)
}

// DeleteProduct ...
func (p *ProductService) DeleteProduct(ctx context.Context, id int) (*models.Product, error) {
	product, err := p.repo.ResolveProductByID(ctx, id)
	if err != nil || product == nil {
		log.Print("product not exist")
		return nil, err
	}
	err = p.repo.DeleteProductByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// GetProductByID ...
func (p *ProductService) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
	return p.repo.ResolveProductByID(ctx, id)
}

// GetAllProducts ...
func (p *ProductService) GetAllProducts(ctx context.Context) ([]*models.Product, error) {
	return p.repo.ResolveProducts(ctx)
}
