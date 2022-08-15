package db

import (
	"fmt"

	"github.com/my-flights/FlightService/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var flights = []model.Flight{
	{
		Model:              gorm.Model{},
		FlightNumber:       "MH526",
		PlaceOfDeparture:   "Belgrade (BEG)",
		PlaceOfArrival:     "Rome (ROM)",
		TimeOfDeparture:    "10/12/2022 9:30 AM",
		TimeOfArrival:      "10/12/2022 11:00 AM",
		AirlineID:          1,
		FlightStatus:       model.ACTIVE,
		EconomyClassPrice:  80,
		BusinessClassPrice: 160,
	},
	{
		Model:              gorm.Model{},
		FlightNumber:       "GH627",
		PlaceOfDeparture:   "Dubai (DXB)",
		PlaceOfArrival:     "Munich (MUC)",
		TimeOfDeparture:    "11/10/2022 4:00 PM",
		TimeOfArrival:      "11/10/2022 11:00 PM",
		AirlineID:          2,
		FlightStatus:       model.ACTIVE,
		EconomyClassPrice:  320,
		BusinessClassPrice: 890,
	},
	{
		Model:              gorm.Model{},
		FlightNumber:       "KL987",
		PlaceOfDeparture:   "Paris (PAR)",
		PlaceOfArrival:     "Barcelona (BCN)",
		TimeOfDeparture:    "8/12/2022 12:00 PM",
		TimeOfArrival:      "8/12/2022 1:00 PM",
		AirlineID:          3,
		FlightStatus:       model.CANCELED,
		EconomyClassPrice:  120,
		BusinessClassPrice: 300,
	},
	{
		Model:              gorm.Model{},
		FlightNumber:       "TS567",
		PlaceOfDeparture:   "Bangkok (BKK)",
		PlaceOfArrival:     "Singapore (SIN)",
		TimeOfDeparture:    "8/10/2022 2:00 PM",
		TimeOfArrival:      "8/10/2022 3:30 PM",
		AirlineID:          4,
		FlightStatus:       model.ACTIVE,
		EconomyClassPrice:  240,
		BusinessClassPrice: 500,
		FirstClassPrice:    2000,
	},
	{
		Model:              gorm.Model{},
		FlightNumber:       "SA679",
		PlaceOfDeparture:   "Belgrade (BEG)",
		PlaceOfArrival:     "Barcelona (BCN)",
		TimeOfDeparture:    "9/11/2022 9:28 AM",
		TimeOfArrival:      "9/11/2022 10:28 AM",
		AirlineID:          5,
		FlightStatus:       model.FULL,
		EconomyClassPrice:  70,
		BusinessClassPrice: 260,
	},
	{
		Model:              gorm.Model{},
		FlightNumber:       "VF231",
		PlaceOfDeparture:   "Belgrade (BEG)",
		PlaceOfArrival:     "Istanbul (IST)",
		TimeOfDeparture:    "8/10/2022 10:00 PM",
		TimeOfArrival:      "8/10/2022 11:30 PM",
		AirlineID:          6,
		FlightStatus:       model.FULL,
		EconomyClassPrice:  180,
		BusinessClassPrice: 410,
	},
	{
		Model:              gorm.Model{},
		FlightNumber:       "TF675",
		PlaceOfDeparture:   "Istanbul (IST)",
		PlaceOfArrival:     "Bangkok (BKK)",
		TimeOfDeparture:    "12/12/2022 1:00 PM",
		TimeOfArrival:      "12/12/2022 11:00 PM",
		AirlineID:          7,
		FlightStatus:       model.ACTIVE,
		EconomyClassPrice:  580,
		BusinessClassPrice: 1200,
		FirstClassPrice:    3500,
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
