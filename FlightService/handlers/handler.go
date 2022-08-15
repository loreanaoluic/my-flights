package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/my-flights/FlightService/repository"
)

type FlightsHandler struct {
	repository *repository.Repository
}

func AdjustResponseHeaderJson(resWriter *http.ResponseWriter) {
	(*resWriter).Header().Set("Content-Type", "application/json")
}

func NewFlightsHandler(repository *repository.Repository) *FlightsHandler {
	return &FlightsHandler{repository}
}

func (rh *FlightsHandler) FindAllFlights(resWriter http.ResponseWriter, req *http.Request) {

	AdjustResponseHeaderJson(&resWriter)

	flightsDTO, _, _ := rh.repository.FindAllFlights(req)

	//json.NewEncoder(resWriter).Encode(model.FlightsPageable{Elements: flights, TotalElements: totalElements})
	json.NewEncoder(resWriter).Encode(flightsDTO)
}
