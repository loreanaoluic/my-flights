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
	router.HandleFunc("/api/flights/update", handler.UpdateFlight).Methods(http.MethodPut)
	router.HandleFunc("/api/flights/cancel/{id}", handler.CancelFlight).Methods(http.MethodPost)
	router.HandleFunc("/api/flights/sort-flights", handler.SortFlights).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8082", router))
}
