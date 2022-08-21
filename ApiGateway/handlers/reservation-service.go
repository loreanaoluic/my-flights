package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/my-flights/ApiGateway/utils"
)

const ReservationsServiceApi string = "/api/reservations"

func FindTicketsByUserId(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	utils.SetupResponse(&w, r)

	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	response, err := http.Get(
		utils.BaseReservationService.Next().Host + ReservationsServiceApi + "/get-all-tickets/" +
			strconv.FormatUint(uint64(id), 10))

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	utils.SetupResponse(&w, r)

	if utils.AuthorizeRole(r, "user") != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	req, _ := http.NewRequest(http.MethodPost,
		utils.BaseReservationService.Next().Host+ReservationsServiceApi+"/book", r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	utils.SetupResponse(&w, r)

	if utils.AuthorizeRole(r, "user") != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodDelete,
		utils.BaseReservationService.Next().Host+ReservationsServiceApi+"/delete/"+strconv.FormatUint(uint64(id), 10),
		r.Body)
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
