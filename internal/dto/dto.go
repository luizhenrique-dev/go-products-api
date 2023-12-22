package dto

import "github.com/luizhenrique-dev/go-products-api/internal/entity"



type ProductCommand struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type CreateProductOutput struct {
	ID string `json:"id"`
}

func (command *ProductCommand) ToEntity() (*entity.Product, error) {
	return entity.NewProduct(command.Name, command.Price, command.Quantity)
}

type GetProductInput struct {
	ID string `json:"id"`
}

type GetProductOutput struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}