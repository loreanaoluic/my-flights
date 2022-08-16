package model

import (
	"github.com/my-flights/AirlineService/repository"
)

func (flight *Flight) ToFlightDTO() FlightDTO {
	airline, _ := repository.FindAirlineById(flight.AirlineID)

	return FlightDTO{
		Id:                 flight.ID,
		FlightNumber:       flight.FlightNumber,
		PlaceOfDeparture:   flight.PlaceOfDeparture,
		PlaceOfArrival:     flight.PlaceOfArrival,
		TimeOfDeparture:    flight.TimeOfDeparture,
		TimeOfArrival:      flight.TimeOfArrival,
		Airline:            airline.Name,
		FlightStatus:       string(flight.FlightStatus),
		EconomyClassPrice:  flight.EconomyClassPrice,
		BusinessClassPrice: flight.BusinessClassPrice,
		FirstClassPrice:    flight.FirstClassPrice,
	}
}
