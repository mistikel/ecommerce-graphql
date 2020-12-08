package product

import (
	"math/rand"
	"time"

	"github.com/graphql-go/graphql"
)

// Handler ...
type Handler struct {
	svc *Service
}

// NewHandler ...
func NewHandler() *Handler {
	return &Handler{
		svc: NewService(),
	}
}

func (h *Handler) getProductHandler(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if !ok {
		return nil, nil
	}
	product, err := h.svc.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (h *Handler) getProductsHandler(p graphql.ResolveParams) (interface{}, error) {
	return h.svc.GetProducts()
}

func (h *Handler) createProductHandler(p graphql.ResolveParams) (interface{}, error) {
	rand.Seed(time.Now().UnixNano())
	product := Product{
		ID:    int64(rand.Intn(100000)), // generate random ID
		Name:  p.Args["name"].(string),
		Price: p.Args["price"].(float64),
	}
	return h.svc.CreateProduct(product)
}

func (h *Handler) updateProductHandler(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(int)
	name, nameOk := p.Args["name"].(string)
	price, priceOk := p.Args["price"].(float64)
	product := Product{
		ID: int64(id),
	}
	if nameOk {
		product.Name = name
	}
	if priceOk {
		product.Price = price
	}

	return h.svc.UpdateProduct(product)
}

func (h *Handler) deleteProductHandler(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(int)
	return h.svc.DeleteProduct(id)
}
