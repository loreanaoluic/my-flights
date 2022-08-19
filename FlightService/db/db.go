package db

import (
	"fmt"

	"github.com/my-flights/FlightService/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var flights = []model.Flight{
	{
		Model:                       gorm.Model{},
		FlightNumber:                "MH526",
		PlaceOfDeparture:            "Belgrade (BEG)",
		PlaceOfArrival:              "Rome (ROM)",
		DateOfDeparture:             "2022-08-18",
		DateOfArrival:               "2022-08-18",
		TimeOfDeparture:             "11:00",
		TimeOfArrival:               "12:00",
		AirlineName:                 "British Airways",
		FlightStatus:                model.ACTIVE,
		EconomyClassPrice:           80,
		BusinessClassPrice:          160,
		EconomyClassRemainingSeats:  5,
		BusinessClassRemainingSeats: 2,
		TimeOfBoarding:              "10:30",
	},
	{
		Model:                       gorm.Model{},
		FlightNumber:                "GH627",
		PlaceOfDeparture:            "Dubai (DXB)",
		PlaceOfArrival:              "Munich (MUC)",
		DateOfDeparture:             "2022-08-18",
		DateOfArrival:               "2022-08-18",
		TimeOfDeparture:             "15:00",
		TimeOfArrival:               "20:00",
		AirlineName:                 "Lufthansa",
		FlightStatus:                model.ACTIVE,
		EconomyClassPrice:           320,
		BusinessClassPrice:          890,
		EconomyClassRemainingSeats:  80,
		BusinessClassRemainingSeats: 20,
		TimeOfBoarding:              "14:30",
	},
	{
		Model:                       gorm.Model{},
		FlightNumber:                "KL987",
		PlaceOfDeparture:            "Paris (PAR)",
		PlaceOfArrival:              "Barcelona (BCN)",
		DateOfDeparture:             "2022-08-20",
		DateOfArrival:               "2022-08-20",
		TimeOfDeparture:             "11:00",
		TimeOfArrival:               "12:00",
		AirlineName:                 "Air France",
		FlightStatus:                model.CANCELED,
		EconomyClassPrice:           120,
		BusinessClassPrice:          300,
		EconomyClassRemainingSeats:  80,
		BusinessClassRemainingSeats: 20,
		TimeOfBoarding:              "10:30",
	},
	{
		Model:                       gorm.Model{},
		FlightNumber:                "TS567",
		PlaceOfDeparture:            "Bangkok (BKK)",
		PlaceOfArrival:              "Singapore (SIN)",
		DateOfDeparture:             "2022-08-22",
		DateOfArrival:               "2022-08-22",
		TimeOfDeparture:             "11:00",
		TimeOfArrival:               "14:00",
		AirlineName:                 "Singapore Airlines",
		FlightStatus:                model.ACTIVE,
		EconomyClassPrice:           240,
		BusinessClassPrice:          500,
		FirstClassPrice:             2000,
		EconomyClassRemainingSeats:  80,
		BusinessClassRemainingSeats: 20,
		FirstClassRemainingSeats:    5,
		TimeOfBoarding:              "10:30",
	},
	{
		Model:                       gorm.Model{},
		FlightNumber:                "SA679",
		PlaceOfDeparture:            "Belgrade (BEG)",
		PlaceOfArrival:              "Barcelona (BCN)",
		DateOfDeparture:             "2022-08-22",
		DateOfArrival:               "2022-08-22",
		TimeOfDeparture:             "11:00",
		TimeOfArrival:               "12:00",
		AirlineName:                 "Lufthansa",
		FlightStatus:                model.FULL,
		EconomyClassPrice:           70,
		BusinessClassPrice:          260,
		EconomyClassRemainingSeats:  80,
		BusinessClassRemainingSeats: 20,
		TimeOfBoarding:              "10:30",
	},
	{
		Model:                       gorm.Model{},
		FlightNumber:                "VF231",
		PlaceOfDeparture:            "Belgrade (BEG)",
		PlaceOfArrival:              "Istanbul (IST)",
		DateOfDeparture:             "2022-08-21",
		DateOfArrival:               "2022-08-21",
		TimeOfDeparture:             "18:00",
		TimeOfArrival:               "19:00",
		AirlineName:                 "Air Serbia",
		FlightStatus:                model.FULL,
		EconomyClassPrice:           180,
		BusinessClassPrice:          410,
		EconomyClassRemainingSeats:  80,
		BusinessClassRemainingSeats: 20,
		TimeOfBoarding:              "17:30",
	},
	{
		Model:                       gorm.Model{},
		FlightNumber:                "TF675",
		PlaceOfDeparture:            "Istanbul (IST)",
		PlaceOfArrival:              "Bangkok (BKK)",
		DateOfDeparture:             "2022-08-25",
		DateOfArrival:               "2022-08-25",
		TimeOfDeparture:             "13:00",
		TimeOfArrival:               "20:00",
		AirlineName:                 "Turkish Airlines",
		FlightStatus:                model.ACTIVE,
		EconomyClassPrice:           580,
		BusinessClassPrice:          1200,
		FirstClassPrice:             3500,
		EconomyClassRemainingSeats:  80,
		BusinessClassRemainingSeats: 20,
		FirstClassRemainingSeats:    5,
		TimeOfBoarding:              "12:30",
	},
}

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=loreana dbname=flights port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to db")
	} else {
		fmt.Println("Database connection successfully created")
	}

	db.Migrator().DropTable("flights")
	db.Migrator().AutoMigrate(&model.Flight{})

	for _, flight := range flights {
		db.Create(&flight)
	}

	return db
}
