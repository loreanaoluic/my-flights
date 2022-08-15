package model

import (
	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	FlightNumber       string       `gorm:"not null;unique"`
	PlaceOfDeparture   string       `gorm:"not null"`
	PlaceOfArrival     string       `gorm:"not null"`
	TimeOfDeparture    string       `gorm:"not null"`
	TimeOfArrival      string       `gorm:"not null"`
	AirlineID          uint         `gorm:"not null"`
	FlightStatus       FlightStatus `gorm:"not null"`
	EconomyClassPrice  float32      `gorm:"not null;min:0.0"`
	BusinessClassPrice float32      `gorm:"not null;min:0.0"`
	FirstClassPrice    float32      `gorm:"min:0.0"`
}
