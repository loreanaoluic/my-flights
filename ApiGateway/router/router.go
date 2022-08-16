package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/my-flights/ApiGateway/handlers"
	"github.com/rs/cors"
)

func HandleRequests() {
	router := mux.NewRouter()

	corsHandler := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"OPTIONS", "HEAD", "GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
	})

	fmt.Println("Listening on: :8080")

	// User Service
	router.HandleFunc("/api/users/login", handlers.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/users/register", handlers.Register).Methods(http.MethodPost)

	// Flight Service
	router.HandleFunc("/api/flights/get-all-flights", handlers.FindAllFlights).Methods(http.MethodGet)

	// Airline Service
	router.HandleFunc("/api/airlines/get-all-airlines", handlers.FindAllAirlines).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", corsHandler.Handler(router)))
}