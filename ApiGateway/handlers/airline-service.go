package handlers

import (
	"net/http"

	"github.com/my-flights/ApiGateway/utils"
)

const AirlinesServiceApi string = "/api/airlines"

func FindAllAirlines(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	utils.SetupResponse(&w, r)

	response, err := http.Get(
		utils.BaseAirlineService.Next().Host + FlightsServiceApi + "/get-all-airlines")

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}
