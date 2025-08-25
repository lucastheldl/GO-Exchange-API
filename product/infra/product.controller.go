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

	product, err := application.CreateProductUseCase(ctx, input, repo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}