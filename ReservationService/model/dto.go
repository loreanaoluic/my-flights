package model

type TicketDTO struct {
	Id               uint    `json:"Id"`
	PlaceOfDeparture string  `gorm:"not null"`
	PlaceOfArrival   string  `gorm:"not null"`
	DateOfDeparture  string  `gorm:"not null"`
	DateOfArrival    string  `gorm:"not null"`
	TimeOfDeparture  string  `gorm:"not null"`
	TimeOfArrival    string  `gorm:"not null"`
	AirlineName      string  `gorm:"not null"`
	Price            float32 `gorm:"min:0.0"`
	TravelClass      string  `gorm:"min:0.0"`
	FlightNumber     string  `gorm:"not null"`
	SeatNumber       string  `gorm:"not null"`
	GateNumber       string  `gorm:"not null"`
	UserId           uint    `gorm:"min:0.0"`
	TimeOfBoarding   string  `gorm:"not null"`
}

type TicketsPageable struct {
	Elements []TicketDTO `json:"Elements"`
	//TotalElements int64    `json:"TotalElements"`
}
