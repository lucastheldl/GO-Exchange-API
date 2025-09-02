package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"go-api/user/application"
	"go-api/user/domain"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

func RegisterRoutes(router *mux.Router, conn *pgx.Conn) {
	handler := &UserHandler{conn: conn}
	router.HandleFunc("/signup", handler.register).Methods("POST")
	router.HandleFunc("/signin", handler.login).Methods("POST")
}

type UserHandler struct {
	conn *pgx.Conn
}


func (h *UserHandler) register(w http.ResponseWriter, r *http.Request){
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
func (h *UserHandler) login(w http.ResponseWriter, r *http.Request){
	var input domain.UserInput


	ctx := context.Background()
	repo := NewUserRepo(h.conn)

	user,err := application.LoginUserUseCase(ctx,input,repo)

	if err != nil{
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}