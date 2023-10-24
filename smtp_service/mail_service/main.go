package main

import (
	"encoding/json"
	"latihan-bottcamp/smtp_service/mail_service/config"
	"log"
	"net/http"

	"github.com/rs/cors"
	"gopkg.in/gomail.v2"
)

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

func sendMailGoMail(from string, to []string, subject, message string, cfg config.Config) (err error) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", from)
	mailer.SetHeader("To", to...)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(
		cfg.Mail.Host,
		cfg.Mail.Port,
		cfg.Mail.Email,
		cfg.Mail.Password,
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}

func sendMail(w http.ResponseWriter, r *http.Request, cfg config.Config) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := sendMailGoMail(req.From, req.To, req.Subject, req.Message, cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := EmailResponse{
		Success: true,
		Message: "Email sent successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	cfg, err := config.LoadConfigYaml("config/config.yaml")
	if err != nil {
		panic(err)
	}

	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	handler := corsOptions.Handler(http.DefaultServeMux)

	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		sendMail(w, r, cfg)
	})

	log.Println("server running at port", cfg.App.Port)
	err = http.ListenAndServe(cfg.App.Port, handler)
	if err != nil {
		panic(err)
	}
}
