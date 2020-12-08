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
	p, err := r.productSvc.CreateProduct(ctx, product.ToProduct(nil))
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *mutationResolver) ProductUpdate(ctx context.Context, id int, product product.Input) (*product.Product, error) {
	p, err := r.productSvc.UpdateProduct(ctx, id, product.ToProduct(&id))
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *mutationResolver) ProductDelete(ctx context.Context, id int) (*product.Product, error) {
	p, err := r.productSvc.DeleteProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productResolver) Variant(ctx context.Context, obj *product.Product) ([]*product.Variant, error) {
	panic("not implemented")
}

func (r *queryResolver) Product(ctx context.Context, id int) (*product.Product, error) {
	p, err := r.productSvc.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *queryResolver) Products(ctx context.Context) ([]product.Product, error) {
	p, err := r.productSvc.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	return p, nil
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
