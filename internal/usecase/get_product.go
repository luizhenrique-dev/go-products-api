package usecase

import (
	"errors"
	"log"

	"github.com/luizhenrique-dev/go-products-api/infra/database"
	"github.com/luizhenrique-dev/go-products-api/internal/dto"
)

var (
	ErrNotFound = errors.New("product not found")
)

type GetProductUC struct {
	ProductRepository database.ProductRepositoryInterface
}

func NewGetProductUC(productRepository database.ProductRepositoryInterface) *GetProductUC {
	return &GetProductUC{
		ProductRepository: productRepository,
	}
}

func (uc *GetProductUC) Execute(input dto.GetProductInput) (*dto.GetProductOutput, error) {
	product, err := uc.ProductRepository.FindById(input.ID)
	if err != nil {
		log.Printf("Error fetching product with id %v", input.ID)
		return nil, ErrNotFound
	}

	log.Printf("Product %v loaded successfully!", product.ID)
	return &dto.GetProductOutput{
		ID:       product.ID.String(),
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}, nil
}
