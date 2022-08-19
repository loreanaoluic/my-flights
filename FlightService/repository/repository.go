package repository

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/my-flights/FlightService/model"

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

func (repo *Repository) FindAllFlights(r *http.Request) ([]model.FlightDTO, int64, error) {
	var flights []model.Flight
	var flightsDTO []model.FlightDTO
	var totalElements int64

	result := repo.db.Scopes(Paginate(r)).Table("flights").Find(&flights)
	repo.db.Table("flights").Count(&totalElements)

	if result.Error != nil {
		return nil, totalElements, result.Error
	}

	for _, flight := range flights {
		flightsDTO = append(flightsDTO, flight.ToFlightDTO())
	}

	return flightsDTO, totalElements, nil
}

func (repo *Repository) SearchFlights(r *http.Request) ([]model.FlightDTO, int64, error) {
	var flightsDTO []model.FlightDTO
	var flights []*model.Flight
	var totalElements int64

	queryParams := r.URL.Query()

	flyingFrom := queryParams.Get("flyingFrom")
	flyingTo := queryParams.Get("flyingTo")
	departing := queryParams.Get("departing")
	passengerNumber, _ := strconv.ParseUint(queryParams.Get("passengerNumber"), 10, 64)
	travelClass, _ := strconv.ParseUint(queryParams.Get("travelClass"), 10, 64)

	if travelClass == 1 {

		result := repo.db.Scopes(Paginate(r)).Table("flights").
			Where("(deleted_at IS NULL) and "+
				"('' = ? or lower(place_of_departure) LIKE ?) and "+
				"('' = ? or lower(place_of_arrival) LIKE ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(economy_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
				flyingTo, "%"+strings.ToLower(flyingTo)+"%",
				departing, "%"+strings.ToLower(departing)+"%",
				passengerNumber,
			).
			Find(&flights)

		repo.db.Table("flights").
			Where("(deleted_at IS NULL) and "+
				"('' = ? or lower(place_of_departure) LIKE ?) and "+
				"('' = ? or lower(place_of_arrival) LIKE ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(economy_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
				flyingTo, "%"+strings.ToLower(flyingTo)+"%",
				departing, "%"+strings.ToLower(departing)+"%",
				passengerNumber,
			).
			Count(&totalElements)

		if result.Error != nil {
			return nil, totalElements, result.Error
		}

	} else if travelClass == 2 {
		result := repo.db.Scopes(Paginate(r)).Table("flights").
			Where("(deleted_at IS NULL) and "+
				"('' = ? or lower(place_of_departure) LIKE ?) and "+
				"('' = ? or lower(place_of_arrival) LIKE ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(business_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
				flyingTo, "%"+strings.ToLower(flyingTo)+"%",
				departing, "%"+strings.ToLower(departing)+"%",
				passengerNumber,
			).
			Find(&flights)

		repo.db.Table("flights").
			Where("(deleted_at IS NULL) and "+
				"('' = ? or lower(place_of_departure) LIKE ?) and "+
				"('' = ? or lower(place_of_arrival) LIKE ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(business_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
				flyingTo, "%"+strings.ToLower(flyingTo)+"%",
				departing, "%"+strings.ToLower(departing)+"%",
				passengerNumber,
			).
			Count(&totalElements)

		if result.Error != nil {
			return nil, totalElements, result.Error
		}

	} else if travelClass == 3 {
		result := repo.db.Scopes(Paginate(r)).Table("flights").
			Where("(deleted_at IS NULL) and "+
				"('' = ? or lower(place_of_departure) LIKE ?) and "+
				"('' = ? or lower(place_of_arrival) LIKE ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(first_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
				flyingTo, "%"+strings.ToLower(flyingTo)+"%",
				departing, "%"+strings.ToLower(departing)+"%",
				passengerNumber,
			).
			Find(&flights)

		repo.db.Table("flights").
			Where("(deleted_at IS NULL) and "+
				"('' = ? or lower(place_of_departure) LIKE ?) and "+
				"('' = ? or lower(place_of_arrival) LIKE ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(first_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
				flyingTo, "%"+strings.ToLower(flyingTo)+"%",
				departing, "%"+strings.ToLower(departing)+"%",
				passengerNumber,
			).
			Count(&totalElements)

		if result.Error != nil {
			return nil, totalElements, result.Error
		}
	}

	for _, flight := range flights {
		flightsDTO = append(flightsDTO, flight.ToFlightDTO())
	}

	return flightsDTO, totalElements, nil
}

func (repo *Repository) CancelFlight(id uint) (*model.FlightDTO, error) {
	var flight model.Flight
	result := repo.db.Table("flights").Where("id = ?", id).First(&flight)

	if result.Error != nil {
		return nil, errors.New("Flight not found!")
	}

	flight.FlightStatus = model.CANCELED

	result2 := repo.db.Table("flights").Save(&flight)

	if result2.Error != nil {
		return nil, errors.New("Flight cannot be canceled!")
	}

	var retValue model.FlightDTO = flight.ToFlightDTO()
	return &retValue, nil
}

func (repo *Repository) CreateFlight(flightDTO *model.FlightDTO) (*model.FlightDTO, error) {
	var flight model.Flight = flightDTO.ToFlight()
	result := repo.db.Table("flights").Create(&flight)

	if result.Error != nil {
		return nil, errors.New("Flight cannot be created!")
	}

	var retValue model.FlightDTO = flight.ToFlightDTO()
	return &retValue, nil
}

func (repo *Repository) UpdateFlight(flightDTO *model.FlightDTO) (*model.FlightDTO, error) {
	var flight model.Flight
	result := repo.db.Table("flights").Where("ID = ?", flightDTO.Id).First(&flight)

	if result.Error != nil {
		return nil, errors.New("Flight cannot be found!")
	}

	flight.FlightNumber = flightDTO.FlightNumber
	flight.PlaceOfDeparture = flightDTO.PlaceOfDeparture
	flight.PlaceOfArrival = flightDTO.PlaceOfArrival
	flight.DateOfDeparture = flightDTO.DateOfDeparture
	flight.DateOfArrival = flightDTO.DateOfArrival
	flight.TimeOfDeparture = flightDTO.TimeOfDeparture
	flight.TimeOfArrival = flightDTO.TimeOfArrival
	flight.AirlineName = flightDTO.Airline
	flight.FlightStatus = model.FlightStatus(flightDTO.FlightStatus)
	flight.EconomyClassPrice = flightDTO.EconomyClassPrice
	flight.BusinessClassPrice = flightDTO.BusinessClassPrice
	flight.FirstClassPrice = flightDTO.FirstClassPrice
	flight.EconomyClassRemainingSeats = flightDTO.EconomyClassRemainingSeats
	flight.BusinessClassRemainingSeats = flightDTO.BusinessClassRemainingSeats
	flight.FirstClassRemainingSeats = flightDTO.FirstClassRemainingSeats

	result2 := repo.db.Table("flights").Save(&flight)

	if result2.Error != nil {
		return nil, errors.New("Flight cannot be updated!")
	}

	var retValue model.FlightDTO = flight.ToFlightDTO()
	return &retValue, nil
}
