package db

import (
	"fmt"

	"github.com/my-flights/AirlineService/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var airlines = []model.Airline{
	{
		Model: gorm.Model{},
		Name:  "Qatar Airways",
	},
	{
		Model: gorm.Model{},
		Name:  "Emirates",
	},
	{
		Model: gorm.Model{},
		Name:  "Air France",
	},
	{
		Model: gorm.Model{},
		Name:  "Lufthansa",
	},
	{
		Model: gorm.Model{},
		Name:  "Turkish Airlines",
	},
	{
		Model: gorm.Model{},
		Name:  "Air Serbia",
	},
	{
		Model: gorm.Model{},
		Name:  "Air Canada",
	},
	{
		Model: gorm.Model{},
		Name:  "British Airways",
	},
	{
		Model: gorm.Model{},
		Name:  "Singapore Airlines",
	},
	{
		Model: gorm.Model{},
		Name:  "Finnair",
	},
}

var Db *gorm.DB

func Init() {
	dsn := "host=localhost user=postgres password=loreana dbname=flights port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to db")
	} else {
		fmt.Println("Database connection successfully created")
	}

	db.Migrator().DropTable("airlines")
	db.Migrator().AutoMigrate(&model.Airline{})

	for _, airline := range airlines {
		db.Create(&airline)
	}
}
