package main

import (
	"ecommerce/server"
)

func main() {
	srv := server.New()
	srv.Serve()
}
