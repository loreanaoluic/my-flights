package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/my-flights/AirlineService/handlers"
)

func HandleRequests(handler *handlers.AirlinesHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/api/airlines/get-all-airlines", handler.FindAllAirlines).Methods(http.MethodGet)
	router.HandleFunc("/api/airlines/create", handler.CreateAirline).Methods(http.MethodPost)
	router.HandleFunc("/api/airlines/update", handler.UpdateAirline).Methods(http.MethodPut)
	router.HandleFunc("/api/airlines/delete/{id}", handler.DeleteAirline).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8083", router))
}
