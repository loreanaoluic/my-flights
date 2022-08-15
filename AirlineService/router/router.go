package router

import (
	"log"
	"net/http"

	"github.com/my-flights/AirlineService/handlers"

	"github.com/gorilla/mux"
)

func HandleRequests(handler *handlers.AirlinesHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/api/airlines/get-all-airlines", handler.FindAllAirlines).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8083", router))
}
