package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/luizhenrique-dev/go-products-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestProductRepository_Create(t *testing.T) {
	// Arrange
	name := "Foo Bar"
	price := 10.50
	quantity := 10
	db := newInMemoryDatabase()
	productRepository := NewProductRepository(db)
	product, _ := entity.NewProduct(name, price, quantity)

	// Act
	err := productRepository.Create(product)

	assert.Nil(t, err)

	var productFromDB entity.Product
	err = db.First(&productFromDB, product.ID).Error

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, productFromDB)
	assert.Equal(t, product.ID, productFromDB.ID)
	assert.Equal(t, product.Name, productFromDB.Name)
	assert.Equal(t, product.Quantity, productFromDB.Quantity)
	assert.Equal(t, product.Price, productFromDB.Price)
}

func TestProductRepository_FindAll(t *testing.T) {
	db := newInMemoryDatabase()
	productRepository := NewProductRepository(db)

	for i := 1; i <= 25; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100, 15)
		assert.NoError(t, err)
		err = productRepository.Create(product)
		assert.NoError(t, err)
	}

	products, err := productRepository.FindAll(1, 10, ASC)
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productRepository.FindAll(2, 10, ASC)
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productRepository.FindAll(3, 10, ASC)
	assert.NoError(t, err)
	assert.Len(t, products, 5)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 25", products[4].Name)
}

func TestProductRepository_FindById(t *testing.T) {
	// Arrange
	name := "Foo Bar"
	price := 10.50
	quantity := 10
	db := newInMemoryDatabase()
	productRepository := NewProductRepository(db)
	product, _ := entity.NewProduct(name, price, quantity)
	err := productRepository.Create(product)
	assert.NoError(t, err)

	// Act
	productFromDB, err := productRepository.FindById(product.ID.String())

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, productFromDB)
	assert.Equal(t, product.ID, productFromDB.ID)
	assert.Equal(t, product.Name, productFromDB.Name)
	assert.Equal(t, product.Quantity, productFromDB.Quantity)
	assert.Equal(t, product.Price, productFromDB.Price)
}

func TestProductRepository_Update(t *testing.T) {
	// Arrange
	db := newInMemoryDatabase()
	productRepository := NewProductRepository(db)
	product, _ := entity.NewProduct("Foo Bar", 10.50, 10)
	err := productRepository.Create(product)
	assert.NoError(t, err)

	// Act
	product.Name = "Bar Foo"
	product.Price = 20.50
	product.Quantity = 20
	err = productRepository.Update(product)

	// Assert
	assert.NoError(t, err)

	var productFromDB entity.Product
	err = db.First(&productFromDB, product.ID).Error

	assert.NoError(t, err)
	assert.NotNil(t, productFromDB)
	assert.Equal(t, product.ID, productFromDB.ID)
	assert.Equal(t, "Bar Foo", productFromDB.Name)
	assert.Equal(t, product.Quantity, productFromDB.Quantity)
	assert.Equal(t, product.Price, productFromDB.Price)
}

func TestProductRepository_Delete(t *testing.T) {
	// Arrange
	db := newInMemoryDatabase()
	productRepository := NewProductRepository(db)
	product, _ := entity.NewProduct("Foo Bar", 10.50, 10)
	err := productRepository.Create(product)
	assert.NoError(t, err)

	// Act
	err = productRepository.Delete(product.ID.String())

	// Assert
	assert.NoError(t, err)

	var productFromDB entity.Product
	err = db.First(&productFromDB, product.ID).Error

	assert.Error(t, err)
	assert.Equal(t, entity.ErrNotFound, err)
}
