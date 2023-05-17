package db

import (
	"fmt"

	"github.com/my-flights/ReservationService/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var tickets = []model.Ticket{
	{
		Model:            gorm.Model{},
		FlightNumber:     "MH526",
		PlaceOfDeparture: "Belgrade (BEG)",
		PlaceOfArrival:   "Rome (ROM)",
		DateOfDeparture:  "2022-08-18",
		DateOfArrival:    "2022-08-18",
		TimeOfDeparture:  "11:00",
		TimeOfArrival:    "12:00",
		AirlineName:      "British Airways",
		Price:            80,
		TravelClass:      model.ECONOMY,
		SeatNumber:       "45A",
		GateNumber:       "11B",
		UserId:           2,
		TimeOfBoarding:   "10:30",
		LosePoints:       2,
	},
	{
		Model:            gorm.Model{},
		FlightNumber:     "GH627",
		PlaceOfDeparture: "Dubai (DXB)",
		PlaceOfArrival:   "Munich (MUC)",
		DateOfDeparture:  "2022-08-18",
		DateOfArrival:    "2022-08-18",
		TimeOfDeparture:  "15:00",
		TimeOfArrival:    "20:00",
		AirlineName:      "Lufthansa",
		Price:            320,
		TravelClass:      model.ECONOMY,
		SeatNumber:       "20C",
		GateNumber:       "10A",
		UserId:           2,
		TimeOfBoarding:   "14:30",
		LosePoints:       2,
	},
	{
		Model:            gorm.Model{},
		FlightNumber:     "KL987",
		PlaceOfDeparture: "Paris (PAR)",
		PlaceOfArrival:   "Barcelona (BCN)",
		DateOfDeparture:  "2022-08-20",
		DateOfArrival:    "2022-08-20",
		TimeOfDeparture:  "11:00",
		TimeOfArrival:    "12:00",
		AirlineName:      "Air France",
		Price:            300,
		TravelClass:      model.BUSINESS,
		SeatNumber:       "2B",
		GateNumber:       "8C",
		UserId:           2,
		TimeOfBoarding:   "10:30",
		LosePoints:       10,
	},
	{
		Model:            gorm.Model{},
		FlightNumber:     "TS567",
		PlaceOfDeparture: "Bangkok (BKK)",
		PlaceOfArrival:   "Singapore (SIN)",
		DateOfDeparture:  "2022-08-22",
		DateOfArrival:    "2022-08-22",
		TimeOfDeparture:  "11:00",
		TimeOfArrival:    "14:00",
		AirlineName:      "Singapore Airlines",
		Price:            2000,
		TravelClass:      model.FIRST,
		SeatNumber:       "1C",
		GateNumber:       "43F",
		UserId:           2,
		TimeOfBoarding:   "10:30",
		LosePoints:       20,
	},
}

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=loreana dbname=flights-reservation-service port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to db")
	} else {
		fmt.Println("Database connection successfully created")
	}

	db.Migrator().DropTable("tickets")
	db.Migrator().AutoMigrate(&model.Ticket{})

	for _, ticket := range tickets {
		db.Create(&ticket)
	}

	return db
}
