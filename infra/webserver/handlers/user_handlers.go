package handlers

import (
	"encoding/json"
	"net/http"
	
	"github.com/luizhenrique-dev/go-products-api/internal/dto"
	usecase "github.com/luizhenrique-dev/go-products-api/internal/usecase/user"
	"github.com/luizhenrique-dev/go-products-api/pkg/security"
)

type UserHandler struct {
	CreateUserUC usecase.CreateUserUC
	GetUserUC    usecase.GetUserUC
	JwtHelper    security.JwtHelper
}

func NewUserHandler(createUserUC usecase.CreateUserUC, getUserUC usecase.GetUserUC, jwtHelper security.JwtHelper) *UserHandler {
	return &UserHandler{
		CreateUserUC: createUserUC,
		GetUserUC:    getUserUC,
		JwtHelper:    jwtHelper,
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

func (handler *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	var inputJwt security.GetJwtInput
	err := json.NewDecoder(r.Body).Decode(&inputJwt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	user, err := handler.GetUserUC.FindByEmail(inputJwt.Email)
	if err != nil && err != usecase.ErrUserNotFound {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Invalid credentials")
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	if !user.ValidatePassword(inputJwt.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Invalid credentials")
		return
	}

	token, err := handler.JwtHelper.GeneratJwt(user.ID.String())
	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: "Bearer " + token,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
}
