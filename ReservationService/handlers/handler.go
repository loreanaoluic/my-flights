package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/my-flights/ReservationService/model"
	"github.com/my-flights/ReservationService/repository"

	"github.com/gorilla/mux"
)

type TicketsHandler struct {
	repository *repository.Repository
}

func NewTicketsHandler(repository *repository.Repository) *TicketsHandler {
	return &TicketsHandler{repository}
}

func AdjustResponseHeaderJson(resWriter *http.ResponseWriter) {
	(*resWriter).Header().Set("Content-Type", "application/json")
}

func (rh *TicketsHandler) FindTicketsByUserId(w http.ResponseWriter, r *http.Request) {

	AdjustResponseHeaderJson(&w)

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.ParseInt(idStr, 10, 64)

	ticketsDTO, _, _ := rh.repository.FindTicketsByUserId(uint(id), r)

	json.NewEncoder(w).Encode(ticketsDTO)
}

func (rh *TicketsHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	var ticketDTO model.TicketDTO
	json.NewDecoder(r.Body).Decode(&ticketDTO)

	createdTicket, err := rh.repository.CreateTicket(ticketDTO.ToTicket())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	} else {
		json.NewEncoder(w).Encode(createdTicket.ToTicketDTO())
	}
}

func (rh *TicketsHandler) DeleteTicket(w http.ResponseWriter, r *http.Request) {
	AdjustResponseHeaderJson(&w)

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.ParseInt(idStr, 10, 64)

	ticketDTO, err := rh.repository.DeleteTicket(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	json.NewEncoder(w).Encode(*ticketDTO)
}
