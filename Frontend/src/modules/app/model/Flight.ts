export interface Flight{
    Id: number
    FlightNumber : string
    PlaceOfDeparture : string
	PlaceOfArrival : string
	DateOfDeparture : string
	DateOfArrival : string
	TimeOfDeparture : string
	TimeOfArrival : string
	Airline : string
	FlightStatus : string
	EconomyClassPrice : number
	BusinessClassPrice : number
	FirstClassPrice : number
	EconomyClassRemainingSeats : number
	BusinessClassRemainingSeats : number
	FirstClassRemainingSeats : number
	TimeOfBoarding : string
}