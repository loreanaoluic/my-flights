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
		TimeOfDeparture:             "2022-08-18 9:30 AM",
		TimeOfArrival:               "2022-08-18 11:00 AM",
		AirlineName:                 "British Airways",
		FlightStatus:                model.ACTIVE,
		EconomyClassPrice:           80,
		BusinessClassPrice:          160,
		EconomyClassRemainingSeats:  80,
		BusinessClassRemainingSeats: 20,
	},
	{
		Model:                       gorm.Model{},
		FlightNumber:                "GH627",
		PlaceOfDeparture:            "Dubai (DXB)",
		PlaceOfArrival:              "Munich (MUC)",
		TimeOfDeparture:             "2022-08-18 4:00 PM",
		TimeOfArrival:               "2022-08-18 11:00 PM",
		AirlineName:                 "Lufthansa",
		FlightStatus:                model.ACTIVE,
		EconomyClassPrice:           320,
		BusinessClassPrice:          890,
		EconomyClassRemainingSeats:  80,
		BusinessClassRemainingSeats: 20,
	},
	{
		Model:                       gorm.Model{},
		FlightNumber:                "KL987",
		PlaceOfDeparture:            "Paris (PAR)",
		PlaceOfArrival:              "Barcelona (BCN)",
		TimeOfDeparture:             "2022-08-18 12:00 PM",
		TimeOfArrival:               "2022-08-18 1:00 PM",
		AirlineName:                 "Air France",
		FlightStatus:                model.CANCELED,
		EconomyClassPrice:           120,
		BusinessClassPrice:          300,
		EconomyClassRemainingSeats:  80,
		BusinessClassRemainingSeats: 20,
	},
	{
		Model:                       gorm.Model{},
		FlightNumber:                "TS567",
		PlaceOfDeparture:            "Bangkok (BKK)",
		PlaceOfArrival:              "Singapore (SIN)",
		TimeOfDeparture:             "2022-08-10 2:00 PM",
		TimeOfArrival:               "2022-08-10 3:30 PM",
		AirlineName:                 "Singapore Airlines",
		FlightStatus:                model.ACTIVE,
		EconomyClassPrice:           240,
		BusinessClassPrice:          500,
		FirstClassPrice:             2000,
		EconomyClassRemainingSeats:  80,
		BusinessClassRemainingSeats: 20,
		FirstClassRemainingSeats:    5,
	},
	{
		Model:                       gorm.Model{},
		FlightNumber:                "SA679",
		PlaceOfDeparture:            "Belgrade (BEG)",
		PlaceOfArrival:              "Barcelona (BCN)",
		TimeOfDeparture:             "2022-09-11 9:28 AM",
		TimeOfArrival:               "2022-09-11 10:28 AM",
		AirlineName:                 "Lufthansa",
		FlightStatus:                model.FULL,
		EconomyClassPrice:           70,
		BusinessClassPrice:          260,
		EconomyClassRemainingSeats:  80,
		BusinessClassRemainingSeats: 20,
	},
	{
		Model:                       gorm.Model{},
		FlightNumber:                "VF231",
		PlaceOfDeparture:            "Belgrade (BEG)",
		PlaceOfArrival:              "Istanbul (IST)",
		TimeOfDeparture:             "2022-08-10 10:00 PM",
		TimeOfArrival:               "2022-08-10 11:30 PM",
		AirlineName:                 "Air Serbia",
		FlightStatus:                model.FULL,
		EconomyClassPrice:           180,
		BusinessClassPrice:          410,
		EconomyClassRemainingSeats:  80,
		BusinessClassRemainingSeats: 20,
	},
	{
		Model:                       gorm.Model{},
		FlightNumber:                "TF675",
		PlaceOfDeparture:            "Istanbul (IST)",
		PlaceOfArrival:              "Bangkok (BKK)",
		TimeOfDeparture:             "2022-08-11 1:00 PM",
		TimeOfArrival:               "2022-08-11 11:00 PM",
		AirlineName:                 "Turkish Airlines",
		FlightStatus:                model.ACTIVE,
		EconomyClassPrice:           580,
		BusinessClassPrice:          1200,
		FirstClassPrice:             3500,
		EconomyClassRemainingSeats:  80,
		BusinessClassRemainingSeats: 20,
		FirstClassRemainingSeats:    5,
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
