package apitets

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/etxnija/productsapi"
	"github.com/etxnija/productsapi/http"
	"github.com/etxnija/productsapi/mysql"
)

func TestGetAll(t *testing.T) {
	repo := mysql.NewMySQLProdutRepository()
	server := http.Server{
		Service: productsapi.NewProductService(repo),
	}
	ts := httptest.NewServer(server.Routes())
	defer ts.Close()

	client := ts.Client()
	resp, err := client.Get(ts.URL + "/api/products")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer resp.Body.Close()
	var products []http.ProductDto
	err = json.NewDecoder(resp.Body).Decode(&products)

	if len(products) != 4 {
		t.Errorf("Should return 4 but got %d", len(products))
	}
}

func TestGetWithSku(t *testing.T) {
	repo := mysql.NewMySQLProdutRepository()
	server := http.Server{
		Service: productsapi.NewProductService(repo),
	}
	ts := httptest.NewServer(server.Routes())
	defer ts.Close()

	client := ts.Client()
	resp, err := client.Get(ts.URL + "/api/products?sku=112233")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer resp.Body.Close()
	var products []http.ProductDto
	err = json.NewDecoder(resp.Body).Decode(&products)

	if len(products) != 1 {
		t.Errorf("Should return 1 but got %d", len(products))
	}
}
