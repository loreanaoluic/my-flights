package handlers

import (
	"net/http"

	"github.com/my-flights/ApiGateway/utils"
)

const UsersServiceApi string = "/api/users"

func Login(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	req, _ := http.NewRequest(http.MethodPost, utils.BaseUserService.Next().Host+UsersServiceApi+"/login", r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func Register(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	req, _ := http.NewRequest(http.MethodPost, utils.BaseUserService.Next().Host+UsersServiceApi+"/register", r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}
