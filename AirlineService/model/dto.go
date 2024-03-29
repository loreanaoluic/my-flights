package model

type AirlineDTO struct {
	Id   uint   `json:"Id"`
	Name string `gorm:"not null;unique"`
}

type ErrorResponse struct {
	Message    string `json:"Message"`
	StatusCode int    `json:"StatusCode"`
}
