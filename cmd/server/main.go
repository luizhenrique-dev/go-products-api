package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
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
	// Recover from panics without crashing server
	r.Use(middleware.Recoverer)

	// Product routes
	r.Route("/products", func(r chi.Router) {
		// Middleware to get the token from the request and set it in the context
		r.Use(jwtauth.Verifier(config.GetTokenAuth()))
		// Middleware to check if the token is valid
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	// User routes
	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJwt)

	http.ListenAndServe(":"+configs.WEB_SERVER_PORT, r)
}
