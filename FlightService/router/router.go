package router

import (
	"log"
	"net/http"

	"github.com/my-flights/FlightService/handlers"

	"github.com/gorilla/mux"
)

func HandleRequests(handler *handlers.FlightsHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/api/flights/get-all-flights", handler.FindAllFlights).Methods(http.MethodGet)
	router.HandleFunc("/api/flights/search-all-flights", handler.SearchFlights).Methods(http.MethodGet)
	router.HandleFunc("/api/flights/create", handler.CreateFlight).Methods(http.MethodPost)
	router.HandleFunc("/api/flights/cancel/{flightNumber}", handler.CancelFlight).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8082", router))
}
