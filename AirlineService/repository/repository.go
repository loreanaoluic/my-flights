package repository

import (
	"net/http"
	"strconv"

	"github.com/my-flights/AirlineService/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

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

func (repo *Repository) FindAllAirlines(r *http.Request) ([]model.Airline, int64, error) {
	var airlines []model.Airline
	var totalElements int64

	result := repo.db.Scopes(Paginate(r)).Table("airlines").Find(&airlines)
	repo.db.Table("airlines").Count(&totalElements)

	if result.Error != nil {
		return nil, totalElements, result.Error
	}

	return airlines, totalElements, nil
}

func (repo *Repository) FindAirlineById(airlineId uint) (model.Airline, error) {
	var airline model.Airline
	repo.db.First(&airline, "ID = ?", airlineId)
	return airline, nil
}
