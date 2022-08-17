package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/my-flights/UserService/model"
	"github.com/my-flights/UserService/repository"
)

type UsersHandler struct {
	repository *repository.Repository
}

func NewUsersHandler(repository *repository.Repository) *UsersHandler {
	return &UsersHandler{repository}
}

func AdjustResponseHeaderJson(resWriter *http.ResponseWriter) {
	(*resWriter).Header().Set("Content-Type", "application/json")
}

func (rh *UsersHandler) Register(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	var userDTO model.RegisterDTO
	json.NewDecoder(r.Body).Decode(&userDTO)

	createdUser, err := rh.repository.Register(userDTO.ToUser())

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	} else {
		json.NewEncoder(w).Encode(createdUser.ToRegisterDTO())
	}
}
