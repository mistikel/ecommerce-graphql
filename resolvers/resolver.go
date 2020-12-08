package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"ecommerce/domain/product"
	"ecommerce/graph/generated"
)

type Resolver struct {
	productSvc *product.Service
}

func (r *mutationResolver) ProductCreate(ctx context.Context, product product.Input) (*product.Product, error) {
	panic("not implemented")
}

func (r *productResolver) Variant(ctx context.Context, obj *product.Product) ([]*product.Variant, error) {
	panic("not implemented")
}

func (r *queryResolver) Product(ctx context.Context, id int) (*product.Product, error) {
	panic("not implemented")
}

func (r *queryResolver) Products(ctx context.Context, limit *int, offset *int) ([]product.Product, error) {
	panic("not implemented")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
