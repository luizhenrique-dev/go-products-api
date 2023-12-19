package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/luizhenrique-dev/go-products-api/internal/dto"
	"github.com/luizhenrique-dev/go-products-api/internal/usecase"
)

type ProductHandler struct {
	CreateProductUC usecase.CreateProductUC
}

func NewProductHandler(createProductUC usecase.CreateProductUC) *ProductHandler {
	return &ProductHandler{
		CreateProductUC: createProductUC,
	}
}

func (handler *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var inputProduct dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&inputProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	
	outputProduct, err := handler.CreateProductUC.Execute(inputProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(outputProduct)
}