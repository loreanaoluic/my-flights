package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	utils.SetupResponse(&w, r)

	response, err := http.Get(
		utils.BaseUserService.Next().Host + UsersServiceApi + "/get-all-users")

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func FindUserById(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	utils.SetupResponse(&w, r)

	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	response, err := http.Get(
		utils.BaseUserService.Next().Host + UsersServiceApi + "/get-one/" + strconv.FormatUint(uint64(id), 10))

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	utils.SetupResponse(&w, r)

	req, _ := http.NewRequest(http.MethodPut,
		utils.BaseUserService.Next().Host+UsersServiceApi+"/update", r.Body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func ActivateAccount(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	utils.SetupResponse(&w, r)

	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodPost,
		utils.BaseUserService.Next().Host+UsersServiceApi+"/activate/"+strconv.FormatUint(uint64(id), 10), r.Body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func DeactivateAccount(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	utils.SetupResponse(&w, r)

	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodPost,
		utils.BaseUserService.Next().Host+UsersServiceApi+"/deactivate/"+strconv.FormatUint(uint64(id), 10), r.Body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func BanUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	utils.SetupResponse(&w, r)

	if utils.AuthorizeRole(r, "admin") != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodPost,
		utils.BaseUserService.Next().Host+UsersServiceApi+"/ban/"+strconv.FormatUint(uint64(id), 10), r.Body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func UnbanUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	utils.SetupResponse(&w, r)

	if utils.AuthorizeRole(r, "admin") != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodPost,
		utils.BaseUserService.Next().Host+UsersServiceApi+"/unban/"+strconv.FormatUint(uint64(id), 10), r.Body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}
