package usecase

import (
	"log"

	"github.com/luizhenrique-dev/go-products-api/infra/database"
	"github.com/luizhenrique-dev/go-products-api/internal/dto"
)

type CreateProductUC struct {
	ProductRepository database.ProductRepositoryInterface
}

func NewCreateProductUC(productRepository database.ProductRepositoryInterface) *CreateProductUC {
	return &CreateProductUC{
		ProductRepository: productRepository,
	}
}

func (uc *CreateProductUC) Execute(input dto.ProductCommand) (*dto.CreateProductOutput, error) {
	product, err := input.ToEntity()
	if err != nil {
		return nil, err
	}

	err = uc.ProductRepository.Create(product)
	if err != nil {
		return nil, err
	}

	log.Printf("Product %v created successfully!", product.ID)
	return &dto.CreateProductOutput{
		ID: product.ID.String(),
	}, nil
}
