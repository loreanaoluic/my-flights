package model

import "gorm.io/gorm"

func (flight *Flight) ToFlightDTO() FlightDTO {

	return FlightDTO{
		Id:                          flight.ID,
		FlightNumber:                flight.FlightNumber,
		PlaceOfDeparture:            flight.PlaceOfDeparture,
		PlaceOfArrival:              flight.PlaceOfArrival,
		DateOfDeparture:             flight.DateOfDeparture,
		DateOfArrival:               flight.DateOfArrival,
		TimeOfDeparture:             flight.TimeOfDeparture,
		TimeOfArrival:               flight.TimeOfArrival,
		Airline:                     flight.AirlineName,
		FlightStatus:                string(flight.FlightStatus),
		EconomyClassPrice:           flight.EconomyClassPrice,
		BusinessClassPrice:          flight.BusinessClassPrice,
		FirstClassPrice:             flight.FirstClassPrice,
		EconomyClassRemainingSeats:  flight.EconomyClassRemainingSeats,
		BusinessClassRemainingSeats: flight.BusinessClassRemainingSeats,
		FirstClassRemainingSeats:    flight.FirstClassRemainingSeats,
		TimeOfBoarding:              flight.TimeOfBoarding,
		EconomyClassPoints:          flight.EconomyClassPoints,
		BusinessClassPoints:         flight.BusinessClassPoints,
		FirstClassPoints:            flight.FirstClassPoints,
		FlightDuration:              flight.FlightDuration,
	}
}

func (flightDTO *FlightDTO) ToFlight() Flight {

	return Flight{
		Model:                       gorm.Model{},
		FlightNumber:                flightDTO.FlightNumber,
		PlaceOfDeparture:            flightDTO.PlaceOfDeparture,
		PlaceOfArrival:              flightDTO.PlaceOfArrival,
		DateOfDeparture:             flightDTO.DateOfDeparture,
		DateOfArrival:               flightDTO.DateOfArrival,
		TimeOfDeparture:             flightDTO.TimeOfDeparture,
		TimeOfArrival:               flightDTO.TimeOfArrival,
		AirlineName:                 flightDTO.Airline,
		FlightStatus:                FlightStatus(flightDTO.FlightStatus),
		EconomyClassPrice:           flightDTO.EconomyClassPrice,
		BusinessClassPrice:          flightDTO.BusinessClassPrice,
		FirstClassPrice:             flightDTO.FirstClassPrice,
		EconomyClassRemainingSeats:  flightDTO.EconomyClassRemainingSeats,
		BusinessClassRemainingSeats: flightDTO.BusinessClassRemainingSeats,
		FirstClassRemainingSeats:    flightDTO.FirstClassRemainingSeats,
		TimeOfBoarding:              flightDTO.TimeOfBoarding,
		EconomyClassPoints:          flightDTO.EconomyClassPoints,
		BusinessClassPoints:         flightDTO.BusinessClassPoints,
		FirstClassPoints:            flightDTO.FirstClassPoints,
		FlightDuration:              flightDTO.FlightDuration,
	}
}
