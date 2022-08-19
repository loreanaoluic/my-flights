package handlers

import (
	"encoding/json"
	"net/http"
	"net/smtp"

	"github.com/gorilla/mux"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email, _ := params["email"]

	w.Header().Set("Content-Type", "application/json")

	from := "myflightswebsite@gmail.com"
	password := "evpszylropemksxs"

	to := []string{
		email,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Subject: Ticket successfully purchased!\n"
	body := "To see detailed flight information, visit the MY TICKETS section on the website."
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}
}
