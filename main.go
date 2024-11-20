package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var redis = make(map[string]int)

// main
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/receipts/process", ProcessHandler).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", GetPointsForIdHandler).Methods("GET")
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
