package usecase

import (
	"log"

	"github.com/luizhenrique-dev/go-products-api/infra/database"
	"github.com/luizhenrique-dev/go-products-api/internal/dto"
)

type CreateUserUC struct {
	UserRepository database.UserRepositoryInterface
}

func NewCreateUserUC(userRepository database.UserRepositoryInterface) *CreateUserUC {
	return &CreateUserUC{
		UserRepository: userRepository,
	}
}

func (uc *CreateUserUC) Execute(input dto.UserCommand) error {
	user, err := input.ToEntity()
	if err != nil {
		return err
	}

	err = uc.UserRepository.Create(user)
	if err != nil {
		return err
	}

	log.Printf("User %v created successfully!", user.ID)
	return nil
}