package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/my-flights/FlightService/model"
	"github.com/my-flights/FlightService/repository"

	"github.com/gorilla/mux"
)

type FlightsHandler struct {
	repository *repository.Repository
}

func NewFlightsHandler(repository *repository.Repository) *FlightsHandler {
	return &FlightsHandler{repository}
}

func AdjustResponseHeaderJson(resWriter *http.ResponseWriter) {
	(*resWriter).Header().Set("Content-Type", "application/json")
}

func (rh *FlightsHandler) FindAllFlights(resWriter http.ResponseWriter, req *http.Request) {

	AdjustResponseHeaderJson(&resWriter)

	flightsDTO, _, _ := rh.repository.FindAllFlights(req)

	json.NewEncoder(resWriter).Encode(flightsDTO)
}

func (rh *FlightsHandler) SearchFlights(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	flightsDTO, _, _ := rh.repository.SearchFlights(req)

	json.NewEncoder(resWriter).Encode(flightsDTO)
}

func (rh *FlightsHandler) CancelFlight(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.ParseInt(idStr, 10, 64)

	flightDTO, err := rh.repository.CancelFlight(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	json.NewEncoder(w).Encode(*flightDTO)
}

func (rh *FlightsHandler) CreateFlight(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	var newFlightDTO model.FlightDTO
	json.NewDecoder(r.Body).Decode(&newFlightDTO)

	flightDTO, err := rh.repository.CreateFlight(&newFlightDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	json.NewEncoder(w).Encode(*flightDTO)
}

func (rh *FlightsHandler) UpdateFlight(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	var updatedFlightDTO model.FlightDTO
	json.NewDecoder(r.Body).Decode(&updatedFlightDTO)

	flightDTO, err := rh.repository.UpdateFlight(&updatedFlightDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	json.NewEncoder(w).Encode(*flightDTO)
}
