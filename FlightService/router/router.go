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

	log.Fatal(http.ListenAndServe(":8082", router))
}