package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"go-api/product/application"
	"go-api/product/domain"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

func RegisterRoutes(router *mux.Router, conn *pgx.Conn) {
	handler := &ProductHandler{conn: conn}
	router.HandleFunc("/products", handler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/:id", handler.EditProduct).Methods("POST")
	router.HandleFunc("/products/:id", handler.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/products", handler.ListAllProducts).Methods("GET")
}

type ProductHandler struct {
	conn *pgx.Conn
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var input domain.ProductInput
	
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	repo := NewProductRepo(h.conn)

	products, err := application.CreateProductUseCase(ctx, input, repo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) ListAllProducts(w http.ResponseWriter, r *http.Request){

	ctx := context.Background()
	repo := NewProductRepo(h.conn)

	products, err := application.GetAllProductsUseCase(ctx, repo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) EditProduct(w http.ResponseWriter, r *http.Request){
	
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request){
	
}

