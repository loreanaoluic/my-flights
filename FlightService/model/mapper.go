package model

func (flight *Flight) ToFlightDTO() FlightDTO {

	return FlightDTO{
		Id:                 flight.ID,
		FlightNumber:       flight.FlightNumber,
		PlaceOfDeparture:   flight.PlaceOfDeparture,
		PlaceOfArrival:     flight.PlaceOfArrival,
		TimeOfDeparture:    flight.TimeOfDeparture,
		TimeOfArrival:      flight.TimeOfArrival,
		Airline:            flight.AirlineName,
		FlightStatus:       string(flight.FlightStatus),
		EconomyClassPrice:  flight.EconomyClassPrice,
		BusinessClassPrice: flight.BusinessClassPrice,
		FirstClassPrice:    flight.FirstClassPrice,
	}
}
