package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	getProductUC := usecase.NewGetProductUC(productRepository)
	updateProductUC := usecase.NewUpdateProductUC(productRepository)
	productHandler := handlers.NewProductHandler(*createProductUC, *getProductUC, *updateProductUC)

	r := chi.NewRouter()
	// Log all requests
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)

	http.ListenAndServe(":"+configs.WEB_SERVER_PORT, r)
}
