package main

import (
	"github.com/my-flights/AirlineService/db"
	"github.com/my-flights/AirlineService/router"
)

func main() {
	db.Init()
	router.HandleRequests()
}
