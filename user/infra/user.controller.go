package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"go-api/user/application"
	"go-api/user/domain"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"

	"net/http"
)

func RegisterRoutes(router *mux.Router, conn *pgx.Conn) {
	handler := &ProductHandler{conn: conn}
	router.HandleFunc("/signup", handler.register).Methods("POST")
}

type ProductHandler struct {
	conn *pgx.Conn
}


func (h *ProductHandler) register(w http.ResponseWriter, r *http.Request){
	var input domain.UserInput


	ctx := context.Background()
	repo := NewUserRepo(h.conn)

	user,err := application.RegisterUserUseCase(ctx,input,repo)

	if err != nil{
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}