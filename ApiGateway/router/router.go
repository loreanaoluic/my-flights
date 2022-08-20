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
	router.HandleFunc("/api/users/get-all-users", handlers.FindAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/users/get-one/{id}", handlers.FindUserById).Methods(http.MethodGet)
	router.HandleFunc("/api/users/update", handlers.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/api/users/ban/{id}", handlers.BanUser).Methods(http.MethodPost)
	router.HandleFunc("/api/users/unban/{id}", handlers.UnbanUser).Methods(http.MethodPost)
	router.HandleFunc("/api/users/activate/{id}", handlers.ActivateAccount).Methods(http.MethodPost)
	router.HandleFunc("/api/users/deactivate/{id}", handlers.DeactivateAccount).Methods(http.MethodPost)
	router.HandleFunc("/api/users/{id}/win/{points}", handlers.WinPoints).Methods(http.MethodPost)
	router.HandleFunc("/api/users/{id}/lose/{points}", handlers.LosePoints).Methods(http.MethodPost)

	// Flight Service
	router.HandleFunc("/api/flights/get-all-flights", handlers.FindAllFlights).Methods(http.MethodGet)
	router.HandleFunc("/api/flights/search-all-flights", handlers.SearchFlights).Methods(http.MethodGet)
	router.HandleFunc("/api/flights/cancel/{id}", handlers.CancelFlight).Methods(http.MethodPost)
	router.HandleFunc("/api/flights/create", handlers.CreateFlight).Methods(http.MethodPost)
	router.HandleFunc("/api/flights/update", handlers.UpdateFlight).Methods(http.MethodPut)

	// Airline Service
	router.HandleFunc("/api/airlines/get-all-airlines", handlers.FindAllAirlines).Methods(http.MethodGet)
	router.HandleFunc("/api/airlines/create", handlers.CreateAirline).Methods(http.MethodPost)
	router.HandleFunc("/api/airlines/update", handlers.UpdateAirline).Methods(http.MethodPut)
	router.HandleFunc("/api/airlines/delete/{id}", handlers.DeleteAirline).Methods(http.MethodDelete)

	// Reservation Service
	router.HandleFunc("/api/reservations/get-all-tickets/{id}", handlers.FindTicketsByUserId).Methods(http.MethodGet)
	router.HandleFunc("/api/reservations/book", handlers.CreateTicket).Methods(http.MethodPost)
	router.HandleFunc("/api/reservations/delete/{id}", handlers.DeleteTicket).Methods(http.MethodDelete)

	// Email Service
	router.HandleFunc("/api/emails/send/{email}", handlers.SendEmail).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", corsHandler.Handler(router)))
}
