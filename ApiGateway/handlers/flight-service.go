package handlers

import (
	"net/http"

	"github.com/my-flights/ApiGateway/utils"
)

const FlightsServiceApi string = "/api/flights"

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
	passengerNumber := queryParams.Get("passengerNumber")
	travelClass := queryParams.Get("travelClass")

	response, err := http.Get(
		utils.BaseFlightService.Next().Host + FlightsServiceApi + "/search-all-flights?flyingFrom=" +
			flyingFrom + "&flyingTo=" + flyingTo + "&departing=" + departing + "&passengerNumber=" + passengerNumber +
			"&travelClass=" + travelClass)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
