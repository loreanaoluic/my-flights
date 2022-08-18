package repository

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/my-flights/AirlineService/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	result := repo.db.Scopes(Paginate(r)).Table("airlines").Where("(deleted_at IS NULL)").Find(&airlines)
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

func (repo *Repository) CreateAirline(airlineDTO *model.AirlineDTO) (*model.AirlineDTO, error) {
	var airline model.Airline = airlineDTO.ToAirline()
	result := repo.db.Table("airlines").Create(&airline)

	if result.Error != nil {
		return nil, errors.New("Airline cannot be created!")
	}

	var retValue model.AirlineDTO = airline.ToAirlineDTO()
	return &retValue, nil
}

func (repo *Repository) UpdateAirline(airlineDTO *model.AirlineDTO) (*model.AirlineDTO, error) {
	var airline model.Airline
	result := repo.db.Table("airlines").Where("ID = ?", airlineDTO.Id).First(&airline)

	if result.Error != nil {
		return nil, errors.New("Airline cannot be found!")
	}

	airline.Name = airlineDTO.Name

	result2 := repo.db.Table("airlines").Save(&airline)

	if result2.Error != nil {
		return nil, errors.New("Airline cannot be updated!")
	}

	var retValue model.AirlineDTO = airline.ToAirlineDTO()
	return &retValue, nil
}

func (repo *Repository) DeleteAirline(id uint) (*model.AirlineDTO, error) {
	var airline model.Airline
	result := repo.db.Table("airlines").Where("id = ?", id).Clauses(clause.Returning{}).Delete(&airline)

	if result.Error != nil {
		return nil, errors.New("Airline cannot be deleted!")
	}

	var retValue model.AirlineDTO = airline.ToAirlineDTO()
	return &retValue, nil
}
