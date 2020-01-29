package http

import (
	"net/http"

	"github.com/etxnija/productsapi"
)

func routes() *http.ServeMux {
	handler := NewProductHandler(productsapi.NewProductService())
	mux := http.NewServeMux()
	mux.HandleFunc("/products", handler.Get())
	return mux
}
