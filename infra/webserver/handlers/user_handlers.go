package handlers

import (
	"encoding/json"
	"net/http"
	
	"github.com/luizhenrique-dev/go-products-api/internal/dto"
	usecase "github.com/luizhenrique-dev/go-products-api/internal/usecase/user"
	"github.com/luizhenrique-dev/go-products-api/pkg/security"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

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

// Create user godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body dto.UserCommand true "user request"
// @Success 201
// @Failure 500	{object} ErrorResponse
// @Router /users [post]
func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var inputUser dto.UserCommand
	err := json.NewDecoder(r.Body).Decode(&inputUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		returnJsonError(err, w, "")
		return
	}

	err = handler.CreateUserUC.Execute(inputUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		returnJsonError(err, w, "")
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetJwt godoc
// @Summary Get a user JWT token
// @Description Get a user JWT token
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body security.GetJwtInput true "user credentials"
// @Success 200 {object} dto.GetJwtOutput
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500	{object} ErrorResponse
// @Router /users/generate_token [post]
func (handler *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	var inputJwt security.GetJwtInput
	err := json.NewDecoder(r.Body).Decode(&inputJwt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		returnJsonError(err, w, "")
		return
	}

	user, err := handler.GetUserUC.FindByEmail(inputJwt.Email)
	if err != nil && err != usecase.ErrUserNotFound {
		w.WriteHeader(http.StatusNotFound)
		returnJsonError(err, w, "Invalid credentials")
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		returnJsonError(err, w, "")
		return
	}

	if !user.ValidatePassword(inputJwt.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		returnJsonError(err, w, "Invalid credentials")
		return
	}

	token, err := handler.JwtHelper.GenerateJwt(user.ID.String())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		returnJsonError(err, w, "")
		return
	}
 
	accessToken := dto.GetJwtOutput{AccessToken: "Bearer " + token}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

func returnJsonError(err error, w http.ResponseWriter, errorMessage string) {
	var errorResponseMessage string
	if errorMessage != "" {
		errorResponseMessage = errorMessage
	} else {
		errorResponseMessage = err.Error()
	}
	errResponse := ErrorResponse{Message: errorResponseMessage}
	json.NewEncoder(w).Encode(errResponse)
}