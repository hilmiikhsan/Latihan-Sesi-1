package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/rs/cors"
)

const BASE_URL = "http://localhost:3000/send"
const PORT = ":4444"

type EmailRequest struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
	Type    string   `json:"type"`
}

type EmailResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func sendMail(w http.ResponseWriter, r *http.Request) {
	var req EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := http.Post(BASE_URL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var emailResp EmailResponse
	if err := json.NewDecoder(resp.Body).Decode(&emailResp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(emailResp)
}

func main() {
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	handler := corsOptions.Handler(http.DefaultServeMux)

	http.HandleFunc("/send", sendMail)

	log.Println("server running at port", PORT)
	err := http.ListenAndServe(PORT, handler)
	if err != nil {
		panic(err)
	}
}
