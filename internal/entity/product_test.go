package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	// Arrange
	name := "Product 1"
	price := 10.0
	quantity := 10

	// Act
	product, err := NewProduct(name, price, quantity)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.CreatedAt)
	assert.Equal(t, name, product.Name)
	assert.Equal(t, price, product.Price)
	assert.Equal(t, quantity, product.Quantity)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	// Arrange
	price := 10.0
	quantity := 10

	// Act
	product, err := NewProduct("", price, quantity)

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	// Arrange
	name := "Product 1"
	quantity := 10

	// Act
	product, err := NewProduct(name, 0, quantity)

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	// Arrange
	name := "Product 1"
	price := -10.0
	quantity := 10

	// Act
	product, err := NewProduct(name, price, quantity)

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductWhenQuantityIsInvalid(t *testing.T) {
	// Arrange
	name := "Product 1"
	price := 10.0
	invalidQuantity := -1

	// Act
	product, err := NewProduct(name, price, invalidQuantity)

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrQuantityIsInvalid, err)
}

func TestProductWhenIdIsRequired(t *testing.T) {
	// Arrange
	name := "Product 1"
	price := 10.0
	quantity := 10

	// Act
	product, _ := NewProduct(name, price, quantity)

	product.ID = uuid.Nil
	err := product.Validate()

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, ErrIdIsRequired, err)
}
