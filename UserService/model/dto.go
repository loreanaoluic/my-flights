package model

import "github.com/dgrijalva/jwt-go"

type RegisterDTO struct {
	Username     string `gorm:"not null;unique"`
	Password     string `gorm:"not null"`
	EmailAddress string `gorm:"not null"`
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
}

type Login struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type Claims struct {
	EmailAddress string   `json:"emailAddress"`
	Role         UserRole `json:"role"`
	Username     string   `json:"username"`
	Id           uint     `json:"Id"`
	jwt.StandardClaims
}

type LoginResponse struct {
	Token string `json:"Token"`
}

type ErrorResponse struct {
	Message    string `json:"Message"`
	StatusCode int    `json:"StatusCode"`
}
