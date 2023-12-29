package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type jsonReq struct {
	Message string `json:"message"`
}

type jsonRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

const (
	port = ":8080"
)

func main() {
	http.HandleFunc("/", handlePostRequest)
	fmt.Printf("Server is listening on port %s...\n", port)
	http.ListenAndServe(port, nil)
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestData jsonReq
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if requestData.Message == "" {
		http.Error(w, "Invalid JSON message", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received message: %s\n", requestData.Message)

	response := jsonRes{
		Status:  "success",
		Message: "Data successfully received",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
