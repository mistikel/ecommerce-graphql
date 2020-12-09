package server

import (
	"ecommerce/database"
	"ecommerce/domain/product"
	"ecommerce/graph/generated"
	"ecommerce/resolvers"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Server ...
type Server struct {
	router     *mux.Router
	productSvc *product.Service
}

// New ...
func New() *Server {
	db, err := gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Couldn't initiate db")
	}
	err = database.Seed(db)
	if err != nil {
		log.Fatal(err)
	}
	s := &Server{
		productSvc: product.NewService(db),
		router:     mux.NewRouter(),
	}
	return s
}

// Serve ...
func (s *Server) Serve() {
	schema := generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolvers.Resolver{
			ProductSvc: s.productSvc,
		},
		Directives: generated.DirectiveRoot{},
		Complexity: generated.ComplexityRoot{},
	})
	srv := handler.NewDefaultServer(schema)
	srv.Use(extension.FixedComplexityLimit(300))
	s.router.Handle("/graphql", srv).Methods("POST")
	s.router.HandleFunc("/", playground.Handler("api-gateway", "/graphql")).Methods("GET")

	color.Green("Listening on localhost:8080\n")
	color.Green("Visit `http://localhost:8080` in your browser\n")
	http.ListenAndServe(":8080", s.router)
}
