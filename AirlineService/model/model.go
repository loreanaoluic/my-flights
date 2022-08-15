package model

import "gorm.io/gorm"

type Airline struct {
	gorm.Model
	Name string `gorm:"not null;unique"`
}

type AirlinesPageable struct {
	Elements []Airline `json:"Elements"`
	//TotalElements int64    `json:"TotalElements"`
}
