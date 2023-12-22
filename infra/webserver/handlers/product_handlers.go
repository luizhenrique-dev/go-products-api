package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/luizhenrique-dev/go-products-api/internal/dto"
	"github.com/luizhenrique-dev/go-products-api/internal/usecase"
)

const EMPTY_STRING = ""

type ProductHandler struct {
	CreateProductUC usecase.CreateProductUC
	GetProductUC    usecase.GetProductUC
	UpdateProductUC usecase.UpdateProductUC
	DeleteProductUC usecase.DeleteProductUC
}

func NewProductHandler(
	createProductUC usecase.CreateProductUC,
	getProductUC usecase.GetProductUC,
	updateProductUC usecase.UpdateProductUC,
	deleteProductUC usecase.DeleteProductUC,
) *ProductHandler {
	return &ProductHandler{
		CreateProductUC: createProductUC,
		GetProductUC:    getProductUC,
		UpdateProductUC: updateProductUC,
		DeleteProductUC: deleteProductUC,
	}
}

func (handler *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var inputProduct dto.ProductCommand
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

func (handler *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == EMPTY_STRING {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("id is required")
		return
	}

	inputGetProduct := dto.GetProductInput{
		ID: id,
	}
	outputGetProduct, err := handler.GetProductUC.Execute(inputGetProduct)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(outputGetProduct)
}

func (handler *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == EMPTY_STRING {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("id is required")
		return
	}

	var command dto.ProductCommand
	err := json.NewDecoder(r.Body).Decode(&command)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = handler.UpdateProductUC.Execute(id, command)
	if err != nil && err.Error() == usecase.ErrNotFound.Error() {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Product updated successfully!")
}

func (handler *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == EMPTY_STRING {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("id is required")
		return
	}

	err := handler.DeleteProductUC.Execute(id)
	if err != nil && err.Error() == usecase.ErrNotFound.Error() {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Product deleted successfully!")
}