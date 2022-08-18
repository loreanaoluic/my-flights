package router

import (
	"log"
	"net/http"

	"github.com/my-flights/UserService/handlers"

	"github.com/gorilla/mux"
)

func HandleRequests(handler *handlers.UsersHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/api/users/login", handler.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/users/register", handler.Register).Methods(http.MethodPost)

	router.HandleFunc("/api/users/authorize/admin", handler.AuthorizeAdmin).Methods(http.MethodGet)
	router.HandleFunc("/api/users/authorize/user", handler.AuthorizeUser).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8081", router))
}
