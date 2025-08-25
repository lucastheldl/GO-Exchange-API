package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// função principal
func main() {
    router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello, World!")
    })
	
	fmt.Println("🚀 Server running at http://localhost:8000")

	log.Fatal(http.ListenAndServe(":8000", router))
}