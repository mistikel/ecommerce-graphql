package server

import (
	"ecommerce/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
)

// Server ...
type Server struct {
	router *mux.Router
	schema *Schema
}

// New ...
func New() *Server {
	db := setupDB()
	s := &Server{
		router: mux.NewRouter(),
		schema: NewSchema(db),
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

func setupDB() *sqlx.DB {
	conf := config.Get()
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", conf.MysqlUsername, conf.MysqlPassword, conf.MysqlHost, conf.MysqlDatabase))
	if err != nil {
		log.Printf("%v", err)
	}
	db.SetMaxOpenConns(conf.MysqlConnectionLimit)
	return db
}
