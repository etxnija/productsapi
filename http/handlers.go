package http

import (
	"encoding/json"
	"log"
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

		var f *productsapi.Filter
		r.ParseForm()
		if len(r.Form) > 0 {
			sku := r.FormValue("sku")
			log.Printf(sku)
			if sku != "" {
				f = &productsapi.Filter{
					Sku: sku,
				}
			}
		}

		products, err := p.productService.GetProducts(productsapi.Page{}, f)

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
