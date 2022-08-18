package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/my-flights/UserService/model"
	"github.com/my-flights/UserService/repository"
)

var jwtKey = []byte("z7031Q8Qy9zVO-T2o7lsFIZSrd05hH0PaeaWIBvLh9s")

type UsersHandler struct {
	repository *repository.Repository
}

func NewUsersHandler(repository *repository.Repository) *UsersHandler {
	return &UsersHandler{repository}
}

func AdjustResponseHeaderJson(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
}

func (uh *UsersHandler) Login(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	var login model.Login
	json.NewDecoder(r.Body).Decode(&login)

	user, err := uh.repository.CheckCredentials(login.Username, login.Password)

	if err != nil {
		if err.Error() == "Invalid username!" || err.Error() == "Invalid password!" {
			w.WriteHeader(http.StatusNotFound)
		} else if err.Error() == "You are banned!" {
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	expirationTime := time.Now().Add(time.Hour * 24)
	claims := model.Claims{EmailAddress: user.EmailAddress, Role: user.Role, Id: user.ID, StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, _ := token.SignedString(jwtKey)

	json.NewEncoder(w).Encode(model.LoginResponse{Token: tokenString})
}

func (rh *UsersHandler) Register(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	var userDTO model.RegisterDTO
	json.NewDecoder(r.Body).Decode(&userDTO)

	createdUser, err := rh.repository.Register(userDTO.ToUser())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	} else {
		json.NewEncoder(w).Encode(createdUser.ToRegisterDTO())
	}
}

func (uh *UsersHandler) AuthorizeAdmin(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	cookie := r.Header.Values("Authorization")
	tokenString := strings.Split(cookie[0], " ")[1]

	claims := model.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if token.Claims.(*model.Claims).Role != model.ADMIN {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func (uh *UsersHandler) AuthorizeUser(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	cookie := r.Header.Values("Authorization")
	tokenString := strings.Split(cookie[0], " ")[1]

	claims := model.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if token.Claims.(*model.Claims).Role != model.USER {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
