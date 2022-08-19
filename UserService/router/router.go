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

	router.HandleFunc("/api/users/get-all-users", handler.FindAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/users/get-one/{id}", handler.FindUserById).Methods(http.MethodGet)
	router.HandleFunc("/api/users/update", handler.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/api/users/ban/{id}", handler.BanUser).Methods(http.MethodPost)
	router.HandleFunc("/api/users/unban/{id}", handler.UnbanUser).Methods(http.MethodPost)
	router.HandleFunc("/api/users/activate/{id}", handler.ActivateAccount).Methods(http.MethodPost)
	router.HandleFunc("/api/users/deactivate/{id}", handler.DeactivateAccount).Methods(http.MethodPost)
	router.HandleFunc("/api/users/{id}/win/{points}", handler.WinPoints).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8081", router))
}
