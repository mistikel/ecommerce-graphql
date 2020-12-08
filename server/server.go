package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
)

// Server ...
type Server struct {
	router *mux.Router
	schema *Schema
}

// New ...
func New() *Server {
	s := &Server{
		router: mux.NewRouter(),
		schema: NewSchema(),
	}
	s.schema.initSchema()
	return s
}

// Serve ...
func (s *Server) Serve() {
	s.router.HandleFunc("/graphql", s.graphQLHandler).Methods("POST")
	http.ListenAndServe(":8080", s.router)
}

func (s *Server) graphQLHandler(w http.ResponseWriter, req *http.Request) {
	var d Data
	if err := json.NewDecoder(req.Body).Decode(&d); err != nil {
		w.WriteHeader(400)
		return
	}
	result := graphql.Do(graphql.Params{
		Context:        req.Context(),
		Schema:         s.schema.Schema,
		RequestString:  d.Query,
		VariableValues: d.Variables,
		OperationName:  d.Operation,
	})
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		fmt.Printf("could not write result to response: %s", err)
	}
}
