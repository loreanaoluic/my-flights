package main

import (
	"github.com/my-flights/FlightService/db"
	"github.com/my-flights/FlightService/handlers"
	"github.com/my-flights/FlightService/repository"
	"github.com/my-flights/FlightService/router"
)

func main() {
	dbConn := db.Init()
	repository := repository.NewRepository(dbConn)
	flightHandler := handlers.NewFlightsHandler(repository)
	router.HandleRequests(flightHandler)
}
