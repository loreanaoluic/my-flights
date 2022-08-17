package model

import "gorm.io/gorm"

type RegisterDTO struct {
	gorm.Model
	Username     string `gorm:"not null;unique"`
	Password     string `gorm:"not null"`
	EmailAddress string `gorm:"not null"`
}

type ErrorResponse struct {
	Message    string `json:"Message"`
	StatusCode int    `json:"StatusCode"`
}
