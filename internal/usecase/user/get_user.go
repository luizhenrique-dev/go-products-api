package usecase

import (
	"errors"
	"log"

	"github.com/luizhenrique-dev/go-products-api/infra/database"
	"github.com/luizhenrique-dev/go-products-api/internal/entity"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type GetUserUC struct {
	UserRepository database.UserRepositoryInterface
}

func NewGetUserUC(userRepository database.UserRepositoryInterface) *GetUserUC {
	return &GetUserUC{
		UserRepository: userRepository,
	}
}

func (uc *GetUserUC) FindByEmail(email string) (*entity.User, error) {
	user, err := uc.UserRepository.FindByEmail(email)
	if err != nil {
		log.Printf("Error fetching user with email %v", email)
		return nil, ErrUserNotFound
	}

	log.Printf("User %v loaded successfully!", user.ID)
	return user, nil
}