package main

import (
	"net/http"

	"github.com/luizhenrique-dev/go-products-api/configs"
	"github.com/luizhenrique-dev/go-products-api/infra/database"
	"github.com/luizhenrique-dev/go-products-api/infra/webserver/handlers"
	"github.com/luizhenrique-dev/go-products-api/internal/usecase"
)

func main() {
	config := configs.NewConfig()
	db := configs.OpenDbConnection(config.GetDBConnectionString())

	productRepository := database.NewProductRepository(db)
	createProductUC := usecase.NewCreateProductUC(productRepository)
	productHandler := handlers.NewProductHandler(*createProductUC)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":"+configs.WEB_SERVER_PORT, nil)
}
