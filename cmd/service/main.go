package main

import (
	"fmt"

	"github.com/etxnija/productsapi"
	"github.com/etxnija/productsapi/http"
	"github.com/etxnija/productsapi/mysql"
)

func main() {
	fmt.Println("This runs")
	repo := mysql.NewMySQLProdutRepository()
	server := http.Server{
		Service: productsapi.NewProductService(repo),
	}
	server.Start()
}
