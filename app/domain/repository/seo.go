package repository

import (
	"context"
	"ecommerce/app/models"
)

// SeoRepository ...
type SeoRepository interface {
	ResolveSeoByID(ctx context.Context, id int) (*models.Seo, error)
	ResolveSeoByProductID(ctx context.Context, id int) (*models.Seo, error)
	ResolveSeos(ctx context.Context) ([]*models.Seo, error)
	StoreSeo(ctx context.Context, seo *models.Seo) error
	UpdateSeo(ctx context.Context, id int, seo *models.Seo) error
	DeleteSeoByID(ctx context.Context, id int) error
}
