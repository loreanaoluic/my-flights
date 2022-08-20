package model

import "gorm.io/gorm"

func (ticket *Ticket) ToTicketDTO() TicketDTO {

	return TicketDTO{
		Id:               ticket.ID,
		PlaceOfDeparture: ticket.PlaceOfDeparture,
		PlaceOfArrival:   ticket.PlaceOfArrival,
		DateOfDeparture:  ticket.DateOfDeparture,
		DateOfArrival:    ticket.DateOfArrival,
		TimeOfDeparture:  ticket.TimeOfDeparture,
		TimeOfArrival:    ticket.TimeOfArrival,
		AirlineName:      ticket.AirlineName,
		Price:            ticket.Price,
		TravelClass:      ticket.TravelClass.String(),
		FlightNumber:     ticket.FlightNumber,
		SeatNumber:       ticket.SeatNumber,
		GateNumber:       ticket.GateNumber,
		UserId:           ticket.UserId,
		TimeOfBoarding:   ticket.TimeOfBoarding,
		LosePoints:       ticket.LosePoints,
	}
}

func (ticketDTO *TicketDTO) ToTicket() Ticket {

	return Ticket{
		Model:            gorm.Model{},
		PlaceOfDeparture: ticketDTO.PlaceOfDeparture,
		PlaceOfArrival:   ticketDTO.PlaceOfArrival,
		DateOfDeparture:  ticketDTO.DateOfDeparture,
		DateOfArrival:    ticketDTO.DateOfArrival,
		TimeOfDeparture:  ticketDTO.TimeOfDeparture,
		TimeOfArrival:    ticketDTO.TimeOfArrival,
		AirlineName:      ticketDTO.AirlineName,
		Price:            ticketDTO.Price,
		TravelClass:      TravelClass(ticketDTO.TravelClass),
		FlightNumber:     ticketDTO.FlightNumber,
		SeatNumber:       ticketDTO.SeatNumber,
		GateNumber:       ticketDTO.GateNumber,
		UserId:           ticketDTO.UserId,
		TimeOfBoarding:   ticketDTO.TimeOfBoarding,
		LosePoints:       ticketDTO.LosePoints,
	}
}
