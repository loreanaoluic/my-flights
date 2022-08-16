package repository

import (
	"net/http"
	"strconv"

	"github.com/my-flights/AirlineService/db"
	"github.com/my-flights/AirlineService/model"

	"gorm.io/gorm"
)

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		if page < 0 {
			page = 0
		}

		pageSize, _ := strconv.Atoi(r.URL.Query().Get("size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := page * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func FindAllAirlines(r *http.Request) ([]model.Airline, int64, error) {
	var airlines []model.Airline
	var totalElements int64

	result := db.Db.Scopes(Paginate(r)).Table("airlines").Find(&airlines)
	db.Db.Table("airlines").Count(&totalElements)

	if result.Error != nil {
		return nil, totalElements, result.Error
	}

	return airlines, totalElements, nil
}

func FindAirlineById(airlineId uint) (model.Airline, error) {
	var airline model.Airline
	db.Db.First(&airline, "ID = ?", airlineId)
	return airline, nil
}
