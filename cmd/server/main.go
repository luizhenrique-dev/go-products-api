package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/luizhenrique-dev/go-products-api/configs"
	"github.com/luizhenrique-dev/go-products-api/infra/database"
	"github.com/luizhenrique-dev/go-products-api/infra/webserver/handlers"
	productUsecase "github.com/luizhenrique-dev/go-products-api/internal/usecase/product"
	userUsecase "github.com/luizhenrique-dev/go-products-api/internal/usecase/user"
	"github.com/luizhenrique-dev/go-products-api/pkg/security"
)

func main() {
	config := configs.NewConfig()
	db := configs.OpenDbConnection(config.GetDBConnectionString())

	// General
	jwtHelper := security.NewJwtHelper(config.GetTokenAuth(), config.GetJwtExpiresIn())

	// Product usecases
	productRepository := database.NewProductRepository(db)
	createProductUC := productUsecase.NewCreateProductUC(productRepository)
	getProductUC := productUsecase.NewGetProductUC(productRepository)
	updateProductUC := productUsecase.NewUpdateProductUC(productRepository)
	deleteProductUC := productUsecase.NewDeleteProductUC(productRepository)

	// User usecases
	userRepository := database.NewUserRepository(db)
	createUserUC := userUsecase.NewCreateUserUC(userRepository)
	getUserUC := userUsecase.NewGetUserUC(userRepository)

	// Handlers
	productHandler := handlers.NewProductHandler(*createProductUC, *getProductUC, *updateProductUC, *deleteProductUC)
	userHandler := handlers.NewUserHandler(*createUserUC, *getUserUC, *jwtHelper)

	r := chi.NewRouter()
	// Log all requests
	r.Use(middleware.Logger)

	// Product routes
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Get("/products", productHandler.GetProducts)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	// User routes
	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJwt)

	http.ListenAndServe(":"+configs.WEB_SERVER_PORT, r)
}
