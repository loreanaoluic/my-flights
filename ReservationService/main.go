package main

import (
	"github.com/my-flights/ReservationService/db"
	"github.com/my-flights/ReservationService/handlers"
	"github.com/my-flights/ReservationService/repository"
	"github.com/my-flights/ReservationService/router"
)

func main() {
	dbConn := db.Init()
	repository := repository.NewRepository(dbConn)
	ticketHandler := handlers.NewTicketsHandler(repository)
	router.HandleRequests(ticketHandler)
}
