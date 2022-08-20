package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"

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
	claims := model.Claims{EmailAddress: user.EmailAddress, Username: user.Username, FirstName: user.FirstName,
		LastName: user.LastName, Role: user.Role, Id: user.ID, Points: user.Points,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}}

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

func (rh *UsersHandler) FindAllUsers(w http.ResponseWriter, r *http.Request) {

	AdjustResponseHeaderJson(&w)

	airlines, _, _ := rh.repository.FindAllUsers(r)

	json.NewEncoder(w).Encode(airlines)
}

func (rh *UsersHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	var updatedUserDTO model.UserDTO
	json.NewDecoder(r.Body).Decode(&updatedUserDTO)

	userDTO, err := rh.repository.UpdateUser(&updatedUserDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	} else {
		json.NewEncoder(w).Encode(*userDTO)
	}
}

func (rh *UsersHandler) FindUserById(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.ParseInt(idStr, 10, 64)

	userDTO, err := rh.repository.FindUserById(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(*userDTO)
}

func (rh *UsersHandler) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.ParseInt(idStr, 10, 64)

	userDTO, err := rh.repository.ActivateAccount(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(*userDTO)
}

func (rh *UsersHandler) DeactivateAccount(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.ParseInt(idStr, 10, 64)

	userDTO, err := rh.repository.DeactivateAccount(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(*userDTO)
}

func (rh *UsersHandler) BanUser(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.ParseInt(idStr, 10, 64)

	userDTO, err := rh.repository.BanUser(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(*userDTO)
}

func (rh *UsersHandler) UnbanUser(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.ParseInt(idStr, 10, 64)

	userDTO, err := rh.repository.UnbanUser(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(*userDTO)
}

func (rh *UsersHandler) WinPoints(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	params := mux.Vars(r)

	userIdStr := params["id"]
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)

	pointsStr := params["points"]
	points, _ := strconv.ParseInt(pointsStr, 10, 64)

	userDTO, err := rh.repository.WinPoints(uint(userId), uint(points))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(*userDTO)
}

func (rh *UsersHandler) LosePoints(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	params := mux.Vars(r)

	userIdStr := params["id"]
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)

	pointsStr := params["points"]
	points, _ := strconv.ParseInt(pointsStr, 10, 64)

	userDTO, err := rh.repository.LosePoints(uint(userId), uint(points))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(*userDTO)
}

func (rh *UsersHandler) BuyTicket(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	params := mux.Vars(r)

	userIdStr := params["id"]
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)

	moneyStr := params["money"]
	money, _ := strconv.ParseFloat(moneyStr, 64)

	userDTO, err := rh.repository.BuyTicket(uint(userId), float64(money))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(*userDTO)
}
