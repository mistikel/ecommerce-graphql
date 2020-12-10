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
	repository.ProductRepository
	repository.VariantRepository
	repository.ShippingRepository
	repository.SeoRepository
}

// NewProductService ...
func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{
		repo.NewProductRepository(db),
		repo.NewVariantRepository(db),
		repo.NewShippingRepository(db),
		repo.NewSeoRepository(db),
	}
}

// CreateProduct ...
func (p *ProductService) CreateProduct(ctx context.Context, input models.Input) (*models.Product, error) {
	product := input.ToProduct()
	err := p.StoreProduct(ctx, product)
	if err != nil {
		return nil, err
	}
	seo := input.Seo.ToSeo(product.ID)
	err = p.StoreSeo(ctx, seo)
	if err != nil {
		return nil, err
	}
	product.Seo = seo
	shipping := input.Shipping.ToShipping(product.ID)
	err = p.StoreShipping(ctx, shipping)
	if err != nil {
		return nil, err
	}
	product.Shipping = shipping
	if !product.IsVariantExist {
		return product, nil
	}
	for _, vi := range input.Variants {
		variant := vi.ToVariant(product.ID)
		err = p.StoreVariant(ctx, variant)
		if err != nil {
			return nil, err
		}
		product.Variant = append(product.Variant, variant)
	}
	return product, nil
}

// UpdateProductByID ...
func (p *ProductService) UpdateProductByID(ctx context.Context, id int, input models.Input) (*models.Product, error) {
	oldProduct, err := p.ResolveProductByID(ctx, id)
	if err != nil || oldProduct == nil {
		log.Print("product not exist")
		return nil, err
	}
	product := input.ToProduct()
	err = p.UpdateProduct(ctx, id, product)
	if err != nil {
		return nil, err
	}
	return p.ResolveProductByID(ctx, product.ID)
}

// DeleteProduct ...
func (p *ProductService) DeleteProduct(ctx context.Context, id int) (*models.Product, error) {
	product, err := p.ResolveProductByID(ctx, id)
	if err != nil || product == nil {
		log.Print("product not exist")
		return nil, err
	}
	err = p.DeleteProductByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// GetProductByID ...
func (p *ProductService) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
	return p.ResolveProductByID(ctx, id)
}

// GetAllProducts ...
func (p *ProductService) GetAllProducts(ctx context.Context) ([]*models.Product, error) {
	return p.ResolveProducts(ctx)
}
