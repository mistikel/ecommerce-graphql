package repository

import (
	"context"
	"ecommerce/app/models"
)

// VariantRepository ...
type VariantRepository interface {
	ResolveVariantByID(ctx context.Context, id int) (*models.Variant, error)
	ResolveVariantByProductID(ctx context.Context, id int) ([]*models.Variant, error)
	ResolveVariants(ctx context.Context) ([]*models.Variant, error)
	StoreVariant(ctx context.Context, variant *models.Variant) error
	UpdateVariant(ctx context.Context, id int, variant *models.Variant) error
	DeleteVariantByID(ctx context.Context, id int) error
}
