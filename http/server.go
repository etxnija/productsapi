package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/etxnija/productsapi"
)

// Server to http
type Server struct {
	Service productsapi.ProductService
}

// Start the server
func (s *Server) Start() {
	// http.HandleFunc("/", hello)
	// http.Handler("/api", routes())

	log.Fatal(http.ListenAndServe(":8080", s.routes()))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
