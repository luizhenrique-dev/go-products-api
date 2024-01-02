package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/luizhenrique-dev/go-products-api/internal/dto"
	usecase "github.com/luizhenrique-dev/go-products-api/internal/usecase/user"
)

type UserHandler struct {
	CreateUserUC usecase.CreateUserUC
}

func NewUserHandler(createUserUC usecase.CreateUserUC) *UserHandler {
	return &UserHandler{
		CreateUserUC: createUserUC,
	}
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var inputUser dto.UserCommand
	err := json.NewDecoder(r.Body).Decode(&inputUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = handler.CreateUserUC.Execute(inputUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}	