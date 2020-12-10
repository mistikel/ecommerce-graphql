package resolvers

import "ecommerce/app/domain/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver ...
type Resolver struct {
	ProductSvc *service.ProductService
}
