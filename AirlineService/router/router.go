package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handler "github.com/my-flights/AirlineService/handlers"
)

func HandleRequests() {
	router := mux.NewRouter()

	router.HandleFunc("/api/airlines/get-all-airlines", handler.FindAllAirlines).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8083", router))
}
