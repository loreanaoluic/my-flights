package model

import "gorm.io/gorm"

func (airline *Airline) ToAirlineDTO() AirlineDTO {

	return AirlineDTO{
		Id:   airline.ID,
		Name: airline.Name,
	}
}

func (airlineDTO *AirlineDTO) ToAirline() Airline {

	return Airline{
		Model: gorm.Model{},
		Name:  airlineDTO.Name,
	}
}
