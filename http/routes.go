package http

import (
	"net/http"
)

func (s *Server) routes() *http.ServeMux {
	handler := NewProductHandler(s.Service)
	mux := http.NewServeMux()
	mux.HandleFunc("/products", handler.Get())
	return mux
}
