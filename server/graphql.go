package server

import (
	"ecommerce/domain/product"

	"github.com/graphql-go/graphql"
)

// Data ...
type Data struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

// Schema ...
type Schema struct {
	Schema        graphql.Schema
	productSchema *product.Schema
}

// NewSchema ...
func NewSchema() *Schema {
	s := &Schema{
		productSchema: product.New(),
	}
	return s
}

func (s *Schema) initSchema() {
	s.Schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    s.initQuery(),
			Mutation: s.initMutation(),
		},
	)
}

func (s *Schema) initQuery() *graphql.Object {
	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Query",
			Fields: make(graphql.Fields),
		},
	)
	for k, v := range s.productSchema.Query() {
		queryType.AddFieldConfig(k, v)
	}
	return queryType
}

func (s *Schema) initMutation() *graphql.Object {
	mutationType := graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Mutation",
			Fields: make(graphql.Fields),
		},
	)
	for k, v := range s.productSchema.Mutation() {
		mutationType.AddFieldConfig(k, v)
	}
	return mutationType
}
