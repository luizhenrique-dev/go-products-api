package usecase

import (
	"log"
	"time"

	"github.com/luizhenrique-dev/go-products-api/infra/database"
	"github.com/luizhenrique-dev/go-products-api/internal/dto"
	"github.com/luizhenrique-dev/go-products-api/internal/entity"
)

type UpdateProductUC struct {
	ProductRepository database.ProductRepositoryInterface
}

func NewUpdateProductUC(productRepository database.ProductRepositoryInterface) *UpdateProductUC {
	return &UpdateProductUC{
		ProductRepository: productRepository,
	}
}

func (uc *UpdateProductUC) Execute(id string, command dto.ProductCommand) error {
	product, err := uc.ProductRepository.FindById(id)
	if err != nil {
		log.Printf("Product with id %v not found!", id)
		return ErrNotFound
	}

	updateProductObject(product, command)
	err = uc.ProductRepository.Update(product)
	if err != nil {
		log.Printf("Error to update product with id %v", id)
		return err
	}

	log.Printf("Product %v updated successfully!", id)
	return nil
}

func updateProductObject(product *entity.Product, command dto.ProductCommand) {
	product.Name = command.Name
	product.Price = command.Price
	product.Quantity = command.Quantity
	product.UpdatedAt = time.Now()
}