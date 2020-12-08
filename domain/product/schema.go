package product

import (
	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
)

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"category": &graphql.Field{
				Type: graphql.String,
			},
			"subCategory": &graphql.Field{
				Type: graphql.String,
			},
			"images": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
			"comparePrice": &graphql.Field{
				Type: graphql.Float,
			},
			"costPerItem": &graphql.Field{
				Type: graphql.Float,
			},
			"shippingCost": &graphql.Field{
				Type: graphql.Float,
			},
			"isChargeTax": &graphql.Field{
				Type: graphql.Boolean,
			},
			"sku": &graphql.Field{
				Type: graphql.String,
			},
			"quantity": &graphql.Field{
				Type: graphql.Int,
			},
			"variant": &graphql.Field{
				Type: variantType,
			},
			"shipping": &graphql.Field{
				Type: shippingType,
			},
			"seo": &graphql.Field{
				Type: seoType,
			},
		},
	},
)

var variantType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Variant",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"skuVariant": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
			"quantuty": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var shippingType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Shipping",
		Fields: graphql.Fields{
			"weight": &graphql.Field{
				Type: graphql.Float,
			},
			"rate": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var seoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Seo",
		Fields: graphql.Fields{
			"auto": &graphql.Field{
				Type: graphql.Boolean,
			},
			"pageTitle": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"tags": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)

// Schema ...
type Schema struct {
	query    graphql.Fields
	mutation graphql.Fields

	handler *Handler
}

// New ...
func New(db *sqlx.DB) *Schema {
	s := &Schema{
		query:    make(graphql.Fields),
		mutation: make(graphql.Fields),
		handler:  NewHandler(db),
	}
	s.setMutation()
	s.setQuery()
	return s
}

// Query get list of query schema
func (s *Schema) Query() graphql.Fields {
	return s.query
}

func (s *Schema) setQuery() {
	s.query = graphql.Fields{
		"product": &graphql.Field{
			Type:        productType,
			Description: "Get product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: s.handler.getProductHandler,
		},

		"list": &graphql.Field{
			Type:        graphql.NewList(productType),
			Description: "Get product list",
			Resolve:     s.handler.getProductsHandler,
		},
	}
}

// Mutation get list of mutation schema
func (s *Schema) Mutation() graphql.Fields {
	return s.mutation
}

func (s *Schema) setMutation() {
	s.mutation = graphql.Fields{
		"create": &graphql.Field{
			Type:        productType,
			Description: "Create new product",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"info": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
			},
			Resolve: s.handler.createProductHandler,
		},
		"update": &graphql.Field{
			Type:        productType,
			Description: "Update product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"info": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.Float,
				},
			},
			Resolve: s.handler.updateProductHandler,
		},
		"delete": &graphql.Field{
			Type:        productType,
			Description: "Delete product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: s.handler.deleteProductHandler,
		},
	}
}
