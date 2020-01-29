package http

import (
	"encoding/json"
	"net/http"

	"github.com/etxnija/productsapi"
)

// ProductHandler have methods to handle product actions
type ProductHandler interface {
	Get() func(w http.ResponseWriter, r *http.Request)
	// GetById(w http.ResponseWriter, r *http.Request)
	// Create(w http.ResponseWriter, r *http.Request)
	// Delete(w http.ResponseWriter, r *http.Request)
	// Update(w http.ResponseWriter, r *http.Request)
}

type productHandler struct {
	productService productsapi.ProductService
}

// NewProductHandler is returned
func NewProductHandler(productService productsapi.ProductService) ProductHandler {
	return &productHandler{
		productService,
	}
}

func (p *productHandler) Get() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		products := []ProductDto{
			ProductDto{
				Id:    "123",
				Title: "Fantastic",
				Sku:   "444555",
			},
			ProductDto{
				Id:    "125",
				Title: "Awesome",
				Sku:   "334444",
			},
		}

		response, err := json.Marshal(products)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
