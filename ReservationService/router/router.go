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
	router.HandleFunc("/api/reservations/get-history/{id}", handler.FindHistoryByUserId).Methods(http.MethodGet)
	router.HandleFunc("/api/reservations/book", handler.CreateTicket).Methods(http.MethodPost)
	router.HandleFunc("/api/reservations/delete/{id}", handler.DeleteTicket).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8084", router))
}
