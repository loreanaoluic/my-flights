package main

import (
	"github.com/my-flights/UserService/db"
	"github.com/my-flights/UserService/handlers"
	"github.com/my-flights/UserService/repository"
	"github.com/my-flights/UserService/router"
)

func main() {
	dbConn := db.Init()
	repository := repository.NewRepository(dbConn)
	userHandler := handlers.NewUsersHandler(repository)
	router.HandleRequests(userHandler)
}
