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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
