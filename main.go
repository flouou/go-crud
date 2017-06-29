package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Printf("Server running")

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/show/{id}", showImageHandle).Methods("GET")
	r.HandleFunc("/save/{id}", saveImageHandle)
}

func indexHandler(w http.ResponseWriter, r *http.Request)