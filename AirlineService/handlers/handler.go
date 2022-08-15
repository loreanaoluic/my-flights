package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/my-flights/AirlineService/repository"
)

type AirlinesHandler struct {
	repository *repository.Repository
}

func AdjustResponseHeaderJson(resWriter *http.ResponseWriter) {
	(*resWriter).Header().Set("Content-Type", "application/json")
}

func NewAirlinesHandler(repository *repository.Repository) *AirlinesHandler {
	return &AirlinesHandler{repository}
}

func (rh *AirlinesHandler) FindAllAirlines(resWriter http.ResponseWriter, req *http.Request) {

	//fmt.Fprint(resWriter, "welcome home")
	AdjustResponseHeaderJson(&resWriter)

	airlines, _, _ := rh.repository.FindAllAirlines(req)

	//json.NewEncoder(resWriter).Encode(model.FlightsPageable{Elements: flights, TotalElements: totalElements})
	json.NewEncoder(resWriter).Encode(airlines)
}
