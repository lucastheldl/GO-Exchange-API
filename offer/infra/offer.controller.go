package infra

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)
type Offerandler struct {
	conn *pgx.Conn
}

func RegisterRoutes(router *mux.Router, conn *pgx.Conn) {
	handler := &Offerandler{conn: conn}
	router.HandleFunc("/offers", handler.ListOffers).Methods("GET")
	router.HandleFunc("/offers", handler.ProposeOffer).Methods("POST")
	router.HandleFunc("/offers/refuse", handler.RefuseOffer).Methods("POST")
	router.HandleFunc("/offers/accept", handler.AcceptOffer).Methods("POST")
}


func (h *Offerandler) ListOffers(w http.ResponseWriter, r *http.Request){

}

func (h *Offerandler) ProposeOffer(w http.ResponseWriter, r *http.Request){

}

func (h *Offerandler) RefuseOffer(w http.ResponseWriter, r *http.Request){

}

func (h *Offerandler) AcceptOffer(w http.ResponseWriter, r *http.Request){

}