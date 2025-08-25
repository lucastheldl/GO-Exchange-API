package infra

import (
	"encoding/json"
	"fmt"
	"go-api/product/application"
	"go-api/product/domain"
	"net/http"
)


func createProduct(w http.ResponseWriter, r *http.Request) {

	var input domain.ProductInput

	err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }


	repo := &ProductRepo{}
	product, err := application.CreateProductUseCase(input,repo)
    if err != nil {
        fmt.Fprint(w, "Error:", err)
        return
    }
    fmt.Fprint(w, product)
}