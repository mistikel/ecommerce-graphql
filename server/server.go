package server

import (
	"net/http"

	"ecommerce/app/domain/service"
	"ecommerce/app/gateway"
	"ecommerce/app/infrastructure/database"
	"ecommerce/app/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

// Server ...
type Server struct {
	router *mux.Router
}

// New ...
func New() *Server {

	s := &Server{
		router: mux.NewRouter(),
	}
	return s
}

func (s *Server) Serve() {
	db := database.OpenDB()
	schema := gateway.NewExecutableSchema(gateway.Config{
		Resolvers: &resolvers.Resolver{
			ProductSvc: service.NewProductService(db),
		},
	})
	srv := handler.NewDefaultServer(schema)
	srv.Use(extension.FixedComplexityLimit(300))
	s.router.Handle("/graphql", srv).Methods("POST")
	s.router.HandleFunc("/", playground.Handler("api-gateway", "/graphql")).Methods("GET")

	color.Green("Listening on localhost:8080\n")
	color.Green("Visit `http://localhost:8080` in your browser\n")
	http.ListenAndServe(":8080", s.router)
}
