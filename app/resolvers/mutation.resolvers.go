package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ecommerce/app/gateway"
	"ecommerce/app/models"
)

func (r *mutationResolver) ProductCreate(ctx context.Context, product models.Input) (*models.Product, error) {
	return r.ProductSvc.CreateProduct(ctx, product)
}

func (r *mutationResolver) ProductUpdate(ctx context.Context, id int, product models.Input) (*models.Product, error) {
	return r.ProductSvc.UpdateProduct(ctx, id, product)
}

func (r *mutationResolver) ProductDelete(ctx context.Context, id int) (*models.Product, error) {
	return r.ProductSvc.DeleteProduct(ctx, id)
}

// Mutation returns gateway.MutationResolver implementation.
func (r *Resolver) Mutation() gateway.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
