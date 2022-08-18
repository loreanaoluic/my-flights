package model

type AirlineDTO struct {
	Id   uint   `json:"Id"`
	Name string `gorm:"not null;unique"`
}
