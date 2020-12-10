package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ecommerce/app/gateway"
	"ecommerce/app/models"
)

func (r *queryResolver) Product(ctx context.Context, id int) (*models.Product, error) {
	return r.ProductSvc.GetProductByID(ctx, id)
}

func (r *queryResolver) Products(ctx context.Context) ([]*models.Product, error) {
	return r.ProductSvc.GetAllProducts(ctx)
}

// Query returns gateway.QueryResolver implementation.
func (r *Resolver) Query() gateway.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
