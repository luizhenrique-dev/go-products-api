package dto

import (
	"time"

	"github.com/luizhenrique-dev/go-products-api/internal/entity"
)

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
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCommand struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (command *UserCommand) ToEntity() (*entity.User, error) {
	return entity.NewUser(command.Name, command.Email, command.Password)
}

type GetJwtOutput struct {
	AccessToken string `json:"access_token"`
}
