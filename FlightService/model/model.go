package model

import (
	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	FlightNumber                string       `gorm:"not null;unique"`
	PlaceOfDeparture            string       `gorm:"not null"`
	PlaceOfArrival              string       `gorm:"not null"`
	DateOfDeparture             string       `gorm:"not null"`
	DateOfArrival               string       `gorm:"not null"`
	TimeOfDeparture             string       `gorm:"not null"`
	TimeOfArrival               string       `gorm:"not null"`
	AirlineName                 string       `gorm:"not null"`
	FlightStatus                FlightStatus `gorm:"not null"`
	EconomyClassPrice           float32      `gorm:"min:0.0"`
	BusinessClassPrice          float32      `gorm:"min:0.0"`
	FirstClassPrice             float32      `gorm:"min:0.0"`
	EconomyClassRemainingSeats  uint         `gorm:"min:0"`
	BusinessClassRemainingSeats uint         `gorm:"min:0"`
	FirstClassRemainingSeats    uint         `gorm:"min:0"`
	TimeOfBoarding              string       `gorm:"not null"`
	EconomyClassPoints          uint         `gorm:"min:0"`
	BusinessClassPoints         uint         `gorm:"min:0"`
	FirstClassPoints            uint         `gorm:"min:0"`
	FlightDuration              uint         `gorm:"min:0"`
}
