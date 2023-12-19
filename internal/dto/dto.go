package dto

import "github.com/luizhenrique-dev/go-products-api/internal/entity"

type CreateProductInput struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
}

type ProductOutput struct {
	ID string `json:"id"`
}

func (input *CreateProductInput) ToEntity() (*entity.Product, error) {
	return entity.NewProduct(input.Name, input.Price, input.Quantity)
}