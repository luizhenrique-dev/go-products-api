package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/luizhenrique-dev/go-products-api/pkg/entity"
)

var (
	ErrIdIsRequired      = errors.New("id is required")
	ErrInvalidId         = errors.New("invalid id")
	ErrNameIsRequired    = errors.New("name is required")
	ErrPriceIsRequired   = errors.New("price is required")
	ErrQuantityIsInvalid = errors.New("invalid quantity")
	ErrInvalidPrice      = errors.New("invalid price")
	ErrNotFound          = errors.New("record not found")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewProduct(name string, price float64, quantity int) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		Quantity:  quantity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID == uuid.Nil {
		return ErrIdIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidId
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price == 0 {
		return ErrPriceIsRequired
	}
	if p.Quantity < 0 {
		return ErrQuantityIsInvalid
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}
