package usecase

import (
	"log"

	"github.com/luizhenrique-dev/go-products-api/infra/database"
)

type DeleteProductUC struct {
	ProductRepository database.ProductRepositoryInterface
}

func NewDeleteProductUC(productRepository database.ProductRepositoryInterface) *DeleteProductUC {
	return &DeleteProductUC{
		ProductRepository: productRepository,
	}
}

func (uc *DeleteProductUC) Execute(id string) error {
	_, err := uc.ProductRepository.FindById(id)
	if err != nil {
		log.Printf("Product with id %v not found!", id)
		return ErrNotFound
	}

	err = uc.ProductRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}