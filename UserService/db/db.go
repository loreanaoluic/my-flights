package db

import (
	"fmt"

	"github.com/my-flights/UserService/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var users = []model.User{
	{
		Model:        gorm.Model{},
		Username:     "admin",
		Password:     "$2a$04$tDyeJl6XBVIxVyMjw6Zau.l1TYP0kZBmpbszswGpJ0j.ScpI4sQ6y",
		EmailAddress: "admin@gmail.com",
		FirstName:    "Milica",
		LastName:     "Markovic",
		Role:         model.UserRole("ADMIN"),
	},
	{
		Model:          gorm.Model{},
		Username:       "user",
		Password:       "$2a$04$CU2TcqokLsIDWIBdOFVN7eoYLUBifthvZhurESow757FeFqpO8FRC",
		EmailAddress:   "fcslrzohkkvzofqkxf@kvhrs.com",
		FirstName:      "Nikola",
		LastName:       "Nikolic",
		Role:           model.UserRole("USER"),
		AccountBalance: 3000,
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

	db.Migrator().DropTable("users")
	db.Migrator().AutoMigrate(&model.User{})

	for _, user := range users {
		db.Create(&user)
	}

	return db
}
