package router

import (
	"log"
	"net/http"

	"github.com/my-flights/ReservationService/handlers"

	"github.com/gorilla/mux"
)

func HandleRequests(handler *handlers.TicketsHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/api/reservations/get-all-tickets/{id}", handler.FindTicketsByUserId).Methods(http.MethodGet)
	router.HandleFunc("/api/reservations/book", handler.CreateTicket).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8084", router))
}