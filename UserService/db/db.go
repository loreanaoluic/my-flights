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
		Password:     "admin",
		EmailAddress: "admin@gmail.com",
	},
	{
		Model:        gorm.Model{},
		Username:     "user",
		Password:     "user",
		EmailAddress: "user@gmail.com",
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
