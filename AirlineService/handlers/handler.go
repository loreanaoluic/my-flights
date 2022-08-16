package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/my-flights/AirlineService/repository"
)

func AdjustResponseHeaderJson(resWriter *http.ResponseWriter) {
	(*resWriter).Header().Set("Content-Type", "application/json")
}

func FindAllAirlines(resWriter http.ResponseWriter, req *http.Request) {

	//fmt.Fprint(resWriter, "welcome home")
	AdjustResponseHeaderJson(&resWriter)

	airlines, _, _ := repository.FindAllAirlines(req)

	//json.NewEncoder(resWriter).Encode(model.FlightsPageable{Elements: flights, TotalElements: totalElements})
	json.NewEncoder(resWriter).Encode(airlines)
}
