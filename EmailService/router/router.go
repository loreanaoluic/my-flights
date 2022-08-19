package router

import (
	"log"
	"net/http"

	handler "github.com/my-flights/EmailService/handlers"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	router := mux.NewRouter()

	router.HandleFunc("/api/emails/send/{email}", handler.SendEmail).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8085", router))
}
