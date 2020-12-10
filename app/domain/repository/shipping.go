package repository

import (
	"context"
	"ecommerce/app/models"
)

// ShippingRepository ...
type ShippingRepository interface {
	ResolveShippingByID(ctx context.Context, id int) (*models.Shipping, error)
	ResolveShippings(ctx context.Context) ([]*models.Shipping, error)
	StoreShipping(ctx context.Context, shipping *models.Shipping) error
	UpdateShipping(ctx context.Context, id int, shipping *models.Shipping) error
	DeleteShippingByID(ctx context.Context, id int) error
}
