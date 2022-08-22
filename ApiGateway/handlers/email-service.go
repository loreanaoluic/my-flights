package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/my-flights/ApiGateway/utils"
)

const EmailsServiceApi string = "/api/emails"

func SendEmail(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	utils.SetupResponse(&w, r)

	if utils.AuthorizeRole(r, "user") != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	email, _ := params["email"]

	req, _ := http.NewRequest(http.MethodPost,
		utils.BaseEmailService.Next().Host+EmailsServiceApi+"/send/"+email, r.Body)
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
