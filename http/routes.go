package http

import (
	"net/http"
)

// Routes for the api
func (s *Server) Routes() *http.ServeMux {
	handler := NewProductHandler(s.Service)
	mux := http.NewServeMux()
	mux.HandleFunc("/api/products", handler.Get())
	return mux
}
