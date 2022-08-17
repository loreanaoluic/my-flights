package main

import (
	"github.com/my-flights/AirlineService/db"
	"github.com/my-flights/AirlineService/handlers"
	"github.com/my-flights/AirlineService/repository"
	"github.com/my-flights/AirlineService/router"
)

func main() {
	dbConn := db.Init()
	repository := repository.NewRepository(dbConn)
	airlineHandler := handlers.NewAirlinesHandler(repository)
	router.HandleRequests(airlineHandler)
}
