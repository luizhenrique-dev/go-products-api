//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/luizhenrique-dev/go-products-api/infra/database"
	productUsecase "github.com/luizhenrique-dev/go-products-api/internal/usecase/product"
	"gorm.io/gorm"
)

var setRepositoryDependency = wire.NewSet(
	database.NewProductRepository,
	wire.Bind(new(database.ProductRepositoryInterface), new(*database.ProductRepository)),
)

func NewCreateProductUseCase(db *gorm.DB) *productUsecase.CreateProductUC {
	wire.Build(
		setRepositoryDependency,
		productUsecase.NewCreateProductUC,
	)
	return &productUsecase.CreateProductUC{}
}
