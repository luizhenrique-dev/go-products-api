package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/luizhenrique-dev/go-products-api/internal/dto"
	usecase "github.com/luizhenrique-dev/go-products-api/internal/usecase/product"
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

// Create product godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body dto.ProductCommand true "product request"
// @Success 201 {object} dto.CreateProductOutput
// @Failure 400	{object} ErrorResponse
// @Failure 500	{object} ErrorResponse
// @Router /products [post]
// @Security ApiKeyAuth
func (handler *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var inputProduct dto.ProductCommand
	err := json.NewDecoder(r.Body).Decode(&inputProduct)
	if err != nil {
		returnJsonError(err, w, "", http.StatusBadRequest)
		return
	}

	outputProduct, err := handler.CreateProductUC.Execute(inputProduct)
	if err != nil {
		returnJsonError(err, w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(outputProduct)
}

// GetProduct godoc
// @Summary Get a product
// @Description Get a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "product id" Format(uuid)
// @Success 200 {object} dto.GetProductOutput
// @Failure 400	{object} ErrorResponse
// @Failure 401
// @Failure 404	{object} ErrorResponse
// @Router /products/{id} [get]
// @Security ApiKeyAuth
func (handler *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == EMPTY_STRING {
		returnJsonError(nil, w, "id is required", http.StatusBadRequest)
		return
	}

	inputGetProduct := dto.GetProductInput{
		ID: id,
	}
	outputGetProduct, err := handler.GetProductUC.Execute(inputGetProduct)
	if err != nil {
		returnJsonError(err, w, "", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(outputGetProduct)
}

// List products godoc
// @Summary List products
// @Description Get all products
// @Tags products
// @Accept  json
// @Produce  json
// @Param page query string false "page number"
// @Param limit query string false "limit per page"
// @Success 200 {array} dto.GetProductOutput
// @Failure 500	{object} ErrorResponse
// @Router /products [get]
// @Security ApiKeyAuth
func (handler *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}

	sort := r.URL.Query().Get("sort")

	products, err := handler.GetProductUC.FindAll(pageInt, limitInt, sort)
	if err != nil {
		returnJsonError(err, w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// Update product godoc
// @Summary Update a product
// @Description Update a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "product id" Format(uuid)
// @Param product body dto.ProductCommand true "product request"
// @Success 200 {string} string "Product updated successfully!"
// @Failure 400	{object} ErrorResponse
// @Failure 401
// @Failure 404	{object} ErrorResponse
// @Failure 500	{object} ErrorResponse
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (handler *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == EMPTY_STRING {
		returnJsonError(nil, w, "id is required", http.StatusBadRequest)
		return
	}

	var command dto.ProductCommand
	err := json.NewDecoder(r.Body).Decode(&command)
	if err != nil {
		returnJsonError(err, w, "", http.StatusBadRequest)
		return
	}

	err = handler.UpdateProductUC.Execute(id, command)
	if err != nil && err.Error() == usecase.ErrNotFound.Error() {
		returnJsonError(err, w, "", http.StatusNotFound)
		return
	}
	if err != nil {
		returnJsonError(err, w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Product updated successfully!")
}

// Delete product godoc
// @Summary Delete a product
// @Description Delete a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "product id" Format(uuid)
// @Success 200 {string} string "Product deleted successfully!"
// @Failure 400	{object} ErrorResponse
// @Failure 401
// @Failure 404	{object} ErrorResponse
// @Failure 500	{object} ErrorResponse
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func (handler *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == EMPTY_STRING {
		returnJsonError(nil, w, "id is required", http.StatusBadRequest)
		return
	}

	err := handler.DeleteProductUC.Execute(id)
	if err != nil && err.Error() == usecase.ErrNotFound.Error() {
		returnJsonError(err, w, "", http.StatusNotFound)
		return
	}
	if err != nil {
		returnJsonError(err, w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Product deleted successfully!")
}
