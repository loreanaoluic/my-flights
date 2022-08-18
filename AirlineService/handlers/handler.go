package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/my-flights/AirlineService/model"
	"github.com/my-flights/AirlineService/repository"
)

type AirlinesHandler struct {
	repository *repository.Repository
}

func NewAirlinesHandler(repository *repository.Repository) *AirlinesHandler {
	return &AirlinesHandler{repository}
}

func AdjustResponseHeaderJson(resWriter *http.ResponseWriter) {
	(*resWriter).Header().Set("Content-Type", "application/json")
}

func (rh *AirlinesHandler) FindAllAirlines(resWriter http.ResponseWriter, req *http.Request) {

	AdjustResponseHeaderJson(&resWriter)

	airlines, _, _ := rh.repository.FindAllAirlines(req)

	//json.NewEncoder(resWriter).Encode(model.FlightsPageable{Elements: flights, TotalElements: totalElements})
	json.NewEncoder(resWriter).Encode(airlines)
}

func (rh *AirlinesHandler) CreateAirline(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	var newAirlineDTO model.AirlineDTO
	json.NewDecoder(r.Body).Decode(&newAirlineDTO)

	airlineDTO, err := rh.repository.CreateAirline(&newAirlineDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(*airlineDTO)
}

func (rh *AirlinesHandler) UpdateAirline(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	var updatedAirlineDTO model.AirlineDTO
	json.NewDecoder(r.Body).Decode(&updatedAirlineDTO)

	airlineDTO, err := rh.repository.UpdateAirline(&updatedAirlineDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(*airlineDTO)
}

func (rh *AirlinesHandler) DeleteAirline(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.ParseInt(idStr, 10, 64)

	airlineDTO, err := rh.repository.DeleteAirline(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(*airlineDTO)
}
