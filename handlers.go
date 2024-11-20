package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	hr := HomeResponse{Message: "Hello there!"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(hr)
}

func ProcessHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data Receipt
	err := decoder.Decode(&data)
	//issue parsing json
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pts := calcTotalPoints(data)
	uuid := uuid.New().String()
	redis[uuid] = pts

	if value, exists := redis[uuid]; exists {
		fmt.Printf("key exists with value: %d\n", value)
	}
	if _, exists := redis["non"]; !exists {
		fmt.Println("does not exists")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	pr := ProcessResponse{Id: uuid}
	json.NewEncoder(w).Encode(pr)

}

func GetPointsForIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if points, exists := redis[id]; exists {
		pr := PointsResponse{Points: points}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(pr)
	} else {
		http.Error(w, "points do not exist for this receipt", http.StatusBadRequest)
	}
}
