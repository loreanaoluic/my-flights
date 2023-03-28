package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/my-flights/ApiGateway/utils"
)

const FlightsServiceApi string = "/api/flights"

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func FindAllFlights(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	utils.SetupResponse(&w, r)

	response, err := http.Get(
		utils.BaseFlightService.Next().Host + FlightsServiceApi + "/get-all-flights")

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func SearchFlights(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	utils.SetupResponse(&w, r)

	queryParams := r.URL.Query()

	flyingFrom := queryParams.Get("flyingFrom")
	flyingTo := queryParams.Get("flyingTo")
	departing := queryParams.Get("departing")
	returning := queryParams.Get("returning")
	passengerNumber := queryParams.Get("passengerNumber")
	travelClass := queryParams.Get("travelClass")
	isReturn := queryParams.Get("isReturn")

	response, err := http.Get(
		utils.BaseFlightService.Next().Host + FlightsServiceApi + "/search-all-flights?flyingFrom=" +
			flyingFrom + "&flyingTo=" + flyingTo + "&departing=" + departing + "&returning=" + returning + "&passengerNumber=" +
			passengerNumber + "&travelClass=" + travelClass + "&isReturn=" + isReturn)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func CancelFlight(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	utils.SetupResponse(&w, r)

	if utils.AuthorizeRole(r, "admin") != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodPost,
		utils.BaseFlightService.Next().Host+FlightsServiceApi+"/cancel/"+strconv.FormatUint(uint64(id), 10), r.Body)
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

func CreateFlight(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	utils.SetupResponse(&w, r)

	if utils.AuthorizeRole(r, "admin") != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	req, _ := http.NewRequest(http.MethodPost,
		utils.BaseFlightService.Next().Host+FlightsServiceApi+"/create", r.Body)
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

func UpdateFlight(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	utils.SetupResponse(&w, r)

	req, _ := http.NewRequest(http.MethodPut,
		utils.BaseFlightService.Next().Host+FlightsServiceApi+"/update", r.Body)
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
