package model

type ErrorResponse struct {
	Message string `json:"message"`
}

type CreateReviewRequest struct {
	UserId    int    `json:"user_id"`
	Username  string `json:"username"`
	AirlineId int    `json:"airline_id"`
	Message   string `json:"message"`
	Rating    int    `json:"rating"`
}
