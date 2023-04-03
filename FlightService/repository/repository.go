package repository

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/my-flights/FlightService/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (repo *Repository) FindAllFlights(r *http.Request) ([]model.FlightDTO, int64, error) {
	var flights []model.Flight
	var flightsDTO []model.FlightDTO
	var totalResults int64

	result := repo.db.Table("flights").Find(&flights)
	repo.db.Table("flights").Count(&totalResults)

	if result.Error != nil {
		return nil, totalResults, result.Error
	}

	for _, flight := range flights {
		flightsDTO = append(flightsDTO, flight.ToFlightDTO())
	}

	return flightsDTO, totalResults, nil
}

func (repo *Repository) SearchFlights(r *http.Request) ([][][]model.FlightDTO, int64, error) {

	flights, totalResults, _ := repo.FindAllByParams(r)

	return flights, totalResults, nil
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

func (repo *Repository) FindAllByParams(r *http.Request) ([][][]model.FlightDTO, int64, error) {
	var allFlights [][][]model.FlightDTO

	var flightsDirect []*model.Flight
	var totalResults int64

	var flights1S [][]*model.Flight
	var flights1SHelp []*model.Flight

	var flights2S [][]*model.Flight
	var flights2SHelp []*model.Flight
	var flights2SHelp2 []*model.Flight

	queryParams := r.URL.Query()

	flyingFrom := queryParams.Get("flyingFrom")
	flyingTo := queryParams.Get("flyingTo")
	departing := queryParams.Get("departing")
	returning := queryParams.Get("returning")
	passengerNumber, _ := strconv.ParseUint(queryParams.Get("passengerNumber"), 10, 64)
	travelClass, _ := strconv.ParseUint(queryParams.Get("travelClass"), 10, 64)
	isReturn := queryParams.Get("isReturn")

	if travelClass == 1 {

		// DIRECT FLIGHTS
		result := repo.db.Table("flights").
			Where("(deleted_at IS NULL) and "+
				"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
				"('' = ? or lower(place_of_arrival) LIKE ? or lower(place_of_arrival) = ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(economy_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%", "%"+strings.ToLower(flyingFrom)+"%",
				flyingTo, "%"+strings.ToLower(flyingTo)+"%", "%"+strings.ToLower(flyingTo)+"%",
				departing, "%"+strings.ToLower(departing)+"%",
				passengerNumber,
			).
			Find(&flightsDirect)

		if result.Error != nil {
			return nil, totalResults, result.Error
		}

		// 1 STOP
		result2 := repo.db.Table("flights").
			Where("(deleted_at IS NULL) and "+
				"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(economy_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%", "%"+strings.ToLower(flyingFrom)+"%",
				departing, "%"+strings.ToLower(departing)+"%",
				passengerNumber,
			).
			Find(&flights1SHelp)

		if result2.Error != nil {
			return nil, totalResults, result.Error
		}

		for _, flight := range flights1SHelp {

			var flightsHelp2 []*model.Flight

			result3 := repo.db.Table("flights").
				Where("(deleted_at IS NULL) and "+
					"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
					"('' = ? or lower(place_of_arrival) LIKE ? or lower(place_of_arrival) = ?) and "+
					"(economy_class_remaining_seats >= ?)",
					flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%", "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%", "%"+strings.ToLower(flyingTo)+"%",
					passengerNumber,
				).
				Find(&flightsHelp2)

			if result3.Error != nil {
				return nil, totalResults, result.Error
			}

			if len(flightsHelp2) != 0 {
				for _, flight2 := range flightsHelp2 {
					var addFlights []*model.Flight

					layout := "2006-01-02 15:04"
					arrival := flight.DateOfArrival + " " + flight.TimeOfArrival
					departure := flight2.DateOfDeparture + " " + flight2.TimeOfDeparture
					arrivalTime, _ := time.Parse(layout, arrival)
					departureTime, _ := time.Parse(layout, departure)

					difference := departureTime.Sub(arrivalTime)

					if arrivalTime.Before(departureTime) {
						if difference.Hours() < 10 {
							addFlights = append(addFlights, flight)
							addFlights = append(addFlights, flight2)
							flights1S = append(flights1S, addFlights)
						}
					}
				}
			}
		}

		// 2 STOPS
		result3 := repo.db.Table("flights").
			Where("(deleted_at IS NULL) and "+
				"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(economy_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%", "%"+strings.ToLower(flyingFrom)+"%",
				departing, "%"+strings.ToLower(departing)+"%",
				passengerNumber,
			).
			Find(&flights2SHelp)

		if result3.Error != nil {
			return nil, totalResults, result.Error
		}

		for _, flight := range flights2SHelp {
			if flight.PlaceOfArrival == flyingTo {
				continue
			}

			result3 := repo.db.Table("flights").
				Where("(deleted_at IS NULL) and "+
					"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
					"(economy_class_remaining_seats >= ?)",
					flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%", "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
					passengerNumber,
				).
				Find(&flights2SHelp2)

			if result3.Error != nil {
				return nil, totalResults, result.Error
			}

			for _, flight2 := range flights2SHelp2 {
				var flightsHelp2 []*model.Flight

				result3 := repo.db.Table("flights").
					Where("(deleted_at IS NULL) and "+
						"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
						"('' = ? or lower(place_of_arrival) LIKE ? or lower(place_of_arrival) = ?) and "+
						"(economy_class_remaining_seats >= ?)",
						flight2.PlaceOfArrival, "%"+strings.ToLower(flight2.PlaceOfArrival)+"%", "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
						flyingTo, "%"+strings.ToLower(flyingTo)+"%", "%"+strings.ToLower(flyingTo)+"%",
						passengerNumber,
					).
					Find(&flightsHelp2)

				if result3.Error != nil {
					return nil, totalResults, result.Error
				}

				if len(flightsHelp2) != 0 {
					for _, flight3 := range flightsHelp2 {
						var addFlights []*model.Flight

						layout := "2006-01-02 15:04"
						flightCalc := flight.DateOfArrival + " " + flight.TimeOfArrival
						flight2Calc := flight2.DateOfDeparture + " " + flight2.TimeOfDeparture
						flight3Calc := flight3.DateOfDeparture + " " + flight3.TimeOfDeparture
						flightTime, _ := time.Parse(layout, flightCalc)
						flight2Time, _ := time.Parse(layout, flight2Calc)
						flight3Time, _ := time.Parse(layout, flight3Calc)

						difference := flight2Time.Sub(flightTime)
						difference2 := flight3Time.Sub(flight2Time)

						if flightTime.Before(flight2Time) && flight2Time.Before(flight3Time) {
							if difference.Hours() < 10 && difference2.Hours() < 10 {
								addFlights = append(addFlights, flight)
								addFlights = append(addFlights, flight2)
								addFlights = append(addFlights, flight3)
								flights2S = append(flights2S, addFlights)
							}
						}
					}
				}
			}
		}

	} else if travelClass == 2 {

		// DIRECT FLIGHTS
		result := repo.db.Table("flights").
			Where("(deleted_at IS NULL) and "+
				"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
				"('' = ? or lower(place_of_arrival) LIKE ? or lower(place_of_arrival) = ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(business_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%", "%"+strings.ToLower(flyingFrom)+"%",
				flyingTo, "%"+strings.ToLower(flyingTo)+"%", "%"+strings.ToLower(flyingTo)+"%",
				departing, "%"+strings.ToLower(departing)+"%",
				passengerNumber,
			).
			Find(&flightsDirect)

		if result.Error != nil {
			return nil, totalResults, result.Error
		}

		// 1 STOP
		result2 := repo.db.Table("flights").
			Where("(deleted_at IS NULL) and "+
				"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(business_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%", "%"+strings.ToLower(flyingFrom)+"%",
				departing, "%"+strings.ToLower(departing)+"%",
				passengerNumber,
			).
			Find(&flights1SHelp)

		if result2.Error != nil {
			return nil, totalResults, result.Error
		}

		for _, flight := range flights1SHelp {

			var flightsHelp2 []*model.Flight

			result3 := repo.db.Table("flights").
				Where("(deleted_at IS NULL) and "+
					"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
					"('' = ? or lower(place_of_arrival) LIKE ? or lower(place_of_arrival) = ?) and "+
					"(business_class_remaining_seats >= ?)",
					flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%", "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%", "%"+strings.ToLower(flyingTo)+"%",
					passengerNumber,
				).
				Find(&flightsHelp2)

			if result3.Error != nil {
				return nil, totalResults, result.Error
			}

			if len(flightsHelp2) != 0 {
				for _, flight2 := range flightsHelp2 {
					var addFlights []*model.Flight

					layout := "2006-01-02 15:04"
					arrival := flight.DateOfArrival + " " + flight.TimeOfArrival
					departure := flight2.DateOfDeparture + " " + flight2.TimeOfDeparture
					arrivalTime, _ := time.Parse(layout, arrival)
					departureTime, _ := time.Parse(layout, departure)

					difference := departureTime.Sub(arrivalTime)

					if arrivalTime.Before(departureTime) {
						if difference.Hours() < 10 {
							addFlights = append(addFlights, flight)
							addFlights = append(addFlights, flight2)
							flights1S = append(flights1S, addFlights)
						}
					}
				}
			}
		}

		// 2 STOPS
		result3 := repo.db.Table("flights").
			Where("(deleted_at IS NULL) and "+
				"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(business_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%", "%"+strings.ToLower(flyingFrom)+"%",
				departing, "%"+strings.ToLower(departing)+"%",
				passengerNumber,
			).
			Find(&flights2SHelp)

		if result3.Error != nil {
			return nil, totalResults, result.Error
		}

		for _, flight := range flights2SHelp {
			if flight.PlaceOfArrival == flyingTo {
				continue
			}
			result3 := repo.db.Table("flights").
				Where("(deleted_at IS NULL) and "+
					"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
					"(business_class_remaining_seats >= ?)",
					flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%", "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
					passengerNumber,
				).
				Find(&flights2SHelp2)

			if result3.Error != nil {
				return nil, totalResults, result.Error
			}

			for _, flight2 := range flights2SHelp2 {
				var flightsHelp2 []*model.Flight

				result3 := repo.db.Table("flights").
					Where("(deleted_at IS NULL) and "+
						"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
						"('' = ? or lower(place_of_arrival) LIKE ? or lower(place_of_arrival) = ?) and "+
						"(business_class_remaining_seats >= ?)",
						flight2.PlaceOfArrival, "%"+strings.ToLower(flight2.PlaceOfArrival)+"%", "%"+strings.ToLower(flight2.PlaceOfArrival)+"%",
						flyingTo, "%"+strings.ToLower(flyingTo)+"%", "%"+strings.ToLower(flyingTo)+"%",
						passengerNumber,
					).
					Find(&flightsHelp2)

				if result3.Error != nil {
					return nil, totalResults, result.Error
				}

				if len(flightsHelp2) != 0 {
					for _, flight3 := range flightsHelp2 {
						var addFlights []*model.Flight

						layout := "2006-01-02 15:04"
						flightCalc := flight.DateOfArrival + " " + flight.TimeOfArrival
						flight2Calc := flight2.DateOfDeparture + " " + flight2.TimeOfDeparture
						flight3Calc := flight3.DateOfDeparture + " " + flight3.TimeOfDeparture
						flightTime, _ := time.Parse(layout, flightCalc)
						flight2Time, _ := time.Parse(layout, flight2Calc)
						flight3Time, _ := time.Parse(layout, flight3Calc)

						if flightTime.Before(flight2Time) && flight2Time.Before(flight3Time) {
							addFlights = append(addFlights, flight)
							addFlights = append(addFlights, flight2)
							addFlights = append(addFlights, flight3)
							flights2S = append(flights2S, addFlights)
						}
					}
				}
			}
		}

	} else if travelClass == 3 {

		// DIRECT FLIGHTS
		result := repo.db.Table("flights").
			Where("(deleted_at IS NULL) and "+
				"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
				"('' = ? or lower(place_of_arrival) LIKE ? or lower(place_of_arrival) = ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(first_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%", "%"+strings.ToLower(flyingFrom)+"%",
				flyingTo, "%"+strings.ToLower(flyingTo)+"%", "%"+strings.ToLower(flyingTo)+"%",
				departing, "%"+strings.ToLower(departing)+"%",
				passengerNumber,
			).
			Find(&flightsDirect)

		if result.Error != nil {
			return nil, totalResults, result.Error
		}

		// 1 STOP
		result2 := repo.db.Table("flights").
			Where("(deleted_at IS NULL) and "+
				"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(first_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%", "%"+strings.ToLower(flyingFrom)+"%",
				departing, "%"+strings.ToLower(departing)+"%",
				passengerNumber,
			).
			Find(&flights1SHelp)

		if result2.Error != nil {
			return nil, totalResults, result.Error
		}

		for _, flight := range flights1SHelp {

			var flightsHelp2 []*model.Flight

			result3 := repo.db.Table("flights").
				Where("(deleted_at IS NULL) and "+
					"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
					"('' = ? or lower(place_of_arrival) LIKE ? or lower(place_of_arrival) = ?) and "+
					"(first_class_remaining_seats >= ?)",
					flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%", "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%", "%"+strings.ToLower(flyingTo)+"%",
					passengerNumber,
				).
				Find(&flightsHelp2)

			if result3.Error != nil {
				return nil, totalResults, result.Error
			}

			if len(flightsHelp2) != 0 {
				for _, flight2 := range flightsHelp2 {
					var addFlights []*model.Flight

					layout := "2006-01-02 15:04"
					arrival := flight.DateOfArrival + " " + flight.TimeOfArrival
					departure := flight2.DateOfDeparture + " " + flight2.TimeOfDeparture
					arrivalTime, _ := time.Parse(layout, arrival)
					departureTime, _ := time.Parse(layout, departure)

					difference := departureTime.Sub(arrivalTime)

					if arrivalTime.Before(departureTime) {
						if difference.Hours() < 10 {
							addFlights = append(addFlights, flight)
							addFlights = append(addFlights, flight2)
							flights1S = append(flights1S, addFlights)
						}
					}
				}
			}
		}

		// 2 STOPS
		result3 := repo.db.Table("flights").
			Where("(deleted_at IS NULL) and "+
				"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(first_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%", "%"+strings.ToLower(flyingFrom)+"%",
				departing, "%"+strings.ToLower(departing)+"%",
				passengerNumber,
			).
			Find(&flights2SHelp)

		if result3.Error != nil {
			return nil, totalResults, result.Error
		}

		for _, flight := range flights2SHelp {
			if flight.PlaceOfArrival == flyingTo {
				continue
			}

			result3 := repo.db.Table("flights").
				Where("(deleted_at IS NULL) and "+
					"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
					"(first_class_remaining_seats >= ?)",
					flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%", "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
					passengerNumber,
				).
				Find(&flights2SHelp2)

			if result3.Error != nil {
				return nil, totalResults, result.Error
			}

			for _, flight2 := range flights2SHelp2 {
				var flightsHelp2 []*model.Flight

				result3 := repo.db.Table("flights").
					Where("(deleted_at IS NULL) and "+
						"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
						"('' = ? or lower(place_of_arrival) LIKE ? or lower(place_of_arrival) = ?) and "+
						"(first_class_remaining_seats >= ?)",
						flight2.PlaceOfArrival, "%"+strings.ToLower(flight2.PlaceOfArrival)+"%", "%"+strings.ToLower(flight2.PlaceOfArrival)+"%",
						flyingTo, "%"+strings.ToLower(flyingTo)+"%", "%"+strings.ToLower(flyingTo)+"%",
						passengerNumber,
					).
					Find(&flightsHelp2)

				if result3.Error != nil {
					return nil, totalResults, result.Error
				}

				if len(flightsHelp2) != 0 {
					for _, flight3 := range flightsHelp2 {
						var addFlights []*model.Flight

						layout := "2006-01-02 15:04"
						flightCalc := flight.DateOfArrival + " " + flight.TimeOfArrival
						flight2Calc := flight2.DateOfDeparture + " " + flight2.TimeOfDeparture
						flight3Calc := flight3.DateOfDeparture + " " + flight3.TimeOfDeparture
						flightTime, _ := time.Parse(layout, flightCalc)
						flight2Time, _ := time.Parse(layout, flight2Calc)
						flight3Time, _ := time.Parse(layout, flight3Calc)

						if flightTime.Before(flight2Time) && flight2Time.Before(flight3Time) {
							addFlights = append(addFlights, flight)
							addFlights = append(addFlights, flight2)
							addFlights = append(addFlights, flight3)
							flights2S = append(flights2S, addFlights)
						}
					}
				}
			}
		}
	}

	for _, flight := range flightsDirect {
		if flight.EconomyClassRemainingSeats == 0 && flight.BusinessClassRemainingSeats == 0 &&
			flight.FirstClassRemainingSeats == 0 {
			flight.FlightStatus = model.FULL
			repo.db.Table("flights").Save(&flight)
		}
	}

	if isReturn == "true" {
		var returnFlightsDirect []*model.Flight
		var returnTotalResults int64

		var returnFlights1S [][]*model.Flight
		var returnFlights1SHelp []*model.Flight

		var returnFlights2S [][]*model.Flight
		var returnFlights2SHelp []*model.Flight
		var returnFlights2SHelp2 []*model.Flight

		if travelClass == 1 {

			// DIRECT FLIGHTS
			result := repo.db.Table("flights").
				Where("(deleted_at IS NULL) and "+
					"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
					"('' = ? or lower(place_of_arrival) LIKE ? or lower(place_of_arrival) = ?) and "+
					"('' = ? or lower(date_of_departure) LIKE ?) and "+
					"(economy_class_remaining_seats >= ?)",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%", "%"+strings.ToLower(flyingTo)+"%",
					flyingFrom, "%"+strings.ToLower(flyingFrom)+"%", "%"+strings.ToLower(flyingFrom)+"%",
					returning, "%"+strings.ToLower(returning)+"%",
					passengerNumber,
				).
				Find(&returnFlightsDirect)

			if result.Error != nil {
				return nil, returnTotalResults, result.Error
			}

			// 1 STOP
			result2 := repo.db.Table("flights").
				Where("(deleted_at IS NULL) and "+
					"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
					"('' = ? or lower(date_of_departure) LIKE ?) and "+
					"(economy_class_remaining_seats >= ?)",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%", "%"+strings.ToLower(flyingTo)+"%",
					returning, "%"+strings.ToLower(returning)+"%",
					passengerNumber,
				).
				Find(&returnFlights1SHelp)

			if result2.Error != nil {
				return nil, returnTotalResults, result.Error
			}

			for _, flight := range returnFlights1SHelp {

				var returnFlightsHelp2 []*model.Flight

				result3 := repo.db.Table("flights").
					Where("(deleted_at IS NULL) and "+
						"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
						"('' = ? or lower(place_of_arrival) LIKE ? or lower(place_of_arrival) = ?) and "+
						"(economy_class_remaining_seats >= ?)",
						flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%", "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
						flyingFrom, "%"+strings.ToLower(flyingFrom)+"%", "%"+strings.ToLower(flyingFrom)+"%",
						passengerNumber,
					).
					Find(&returnFlightsHelp2)

				if result3.Error != nil {
					return nil, returnTotalResults, result.Error
				}

				if len(returnFlightsHelp2) != 0 {
					for _, flight2 := range returnFlightsHelp2 {

						var addReturnFlights []*model.Flight

						layout := "2006-01-02 15:04"
						arrival := flight.DateOfArrival + " " + flight.TimeOfArrival
						departure := flight2.DateOfDeparture + " " + flight2.TimeOfDeparture
						arrivalTime, _ := time.Parse(layout, arrival)
						departureTime, _ := time.Parse(layout, departure)

						difference := departureTime.Sub(arrivalTime)

						if arrivalTime.Before(departureTime) {
							if difference.Hours() < 10 {
								addReturnFlights = append(addReturnFlights, flight)
								addReturnFlights = append(addReturnFlights, flight2)
								returnFlights1S = append(returnFlights1S, addReturnFlights)
							}
						}
					}
				}
			}

			// 2 STOPS
			result3 := repo.db.Table("flights").
				Where("(deleted_at IS NULL) and "+
					"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
					"('' = ? or lower(date_of_departure) LIKE ?) and "+
					"(economy_class_remaining_seats >= ?)",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%", "%"+strings.ToLower(flyingTo)+"%",
					returning, "%"+strings.ToLower(returning)+"%",
					passengerNumber,
				).
				Find(&returnFlights2SHelp)

			if result3.Error != nil {
				return nil, totalResults, result.Error
			}

			for _, flight := range returnFlights2SHelp {

				result3 := repo.db.Table("flights").
					Where("(deleted_at IS NULL) and "+
						"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
						"(economy_class_remaining_seats >= ?)",
						flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%", "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
						passengerNumber,
					).
					Find(&returnFlights2SHelp2)

				if result3.Error != nil {
					return nil, totalResults, result.Error
				}

				for _, flight2 := range returnFlights2SHelp2 {
					if flight2.PlaceOfArrival == flyingTo {
						continue
					}

					var flightsHelp2 []*model.Flight

					result3 := repo.db.Table("flights").
						Where("(deleted_at IS NULL) and "+
							"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
							"('' = ? or lower(place_of_arrival) LIKE ? or lower(place_of_arrival) = ?) and "+
							"(economy_class_remaining_seats >= ?)",
							flight2.PlaceOfArrival, "%"+strings.ToLower(flight2.PlaceOfArrival)+"%", "%"+strings.ToLower(flight2.PlaceOfArrival)+"%",
							flyingFrom, "%"+strings.ToLower(flyingFrom)+"%", "%"+strings.ToLower(flyingFrom)+"%",
							passengerNumber,
						).
						Find(&flightsHelp2)

					if result3.Error != nil {
						return nil, totalResults, result.Error
					}

					if len(flightsHelp2) != 0 {
						for _, flight3 := range flightsHelp2 {
							var addFlights []*model.Flight

							layout := "2006-01-02 15:04"
							flightCalc := flight.DateOfArrival + " " + flight.TimeOfArrival
							flight2Calc := flight2.DateOfDeparture + " " + flight2.TimeOfDeparture
							flight3Calc := flight3.DateOfDeparture + " " + flight3.TimeOfDeparture
							flightTime, _ := time.Parse(layout, flightCalc)
							flight2Time, _ := time.Parse(layout, flight2Calc)
							flight3Time, _ := time.Parse(layout, flight3Calc)

							difference1 := flight2Time.Sub(flightTime)
							difference2 := flight3Time.Sub(flight2Time)

							if flightTime.Before(flight2Time) && flight2Time.Before(flight3Time) {
								if difference1.Hours() < 10 && difference2.Hours() < 10 {
									addFlights = append(addFlights, flight)
									addFlights = append(addFlights, flight2)
									addFlights = append(addFlights, flight3)
									returnFlights2S = append(returnFlights2S, addFlights)
								}
							}
						}
					}
				}
			}

		} else if travelClass == 2 {

			// DIRECT FLIGHTS
			result := repo.db.Table("flights").
				Where("(deleted_at IS NULL) and "+
					"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
					"('' = ? or lower(place_of_arrival) LIKE ? or lower(place_of_arrival) = ?) and "+
					"('' = ? or lower(date_of_departure) LIKE ?) and "+
					"(business_class_remaining_seats >= ?)",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%", "%"+strings.ToLower(flyingTo)+"%",
					flyingFrom, "%"+strings.ToLower(flyingFrom)+"%", "%"+strings.ToLower(flyingFrom)+"%",
					returning, "%"+strings.ToLower(returning)+"%",
					passengerNumber,
				).
				Find(&returnFlightsDirect)

			if result.Error != nil {
				return nil, returnTotalResults, result.Error
			}

			// 1 STOP
			result2 := repo.db.Table("flights").
				Where("(deleted_at IS NULL) and "+
					"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
					"('' = ? or lower(date_of_departure) LIKE ?) and "+
					"(business_class_remaining_seats >= ?)",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%", "%"+strings.ToLower(flyingTo)+"%",
					returning, "%"+strings.ToLower(returning)+"%",
					passengerNumber,
				).
				Find(&returnFlights1SHelp)

			if result2.Error != nil {
				return nil, returnTotalResults, result.Error
			}

			for _, flight := range returnFlights1SHelp {

				var returnFlightsHelp2 []*model.Flight

				result3 := repo.db.Table("flights").
					Where("(deleted_at IS NULL) and "+
						"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
						"('' = ? or lower(place_of_arrival) LIKE ? or lower(place_of_arrival) = ?) and "+
						"(business_class_remaining_seats >= ?)",
						flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%", "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
						flyingFrom, "%"+strings.ToLower(flyingFrom)+"%", "%"+strings.ToLower(flyingFrom)+"%",
						passengerNumber,
					).
					Find(&returnFlightsHelp2)

				if result3.Error != nil {
					return nil, returnTotalResults, result.Error
				}

				if len(returnFlightsHelp2) != 0 {
					for _, flight2 := range returnFlightsHelp2 {

						var addReturnFlights []*model.Flight

						layout := "2006-01-02 15:04"
						arrival := flight.DateOfArrival + " " + flight.TimeOfArrival
						departure := flight2.DateOfDeparture + " " + flight2.TimeOfDeparture
						arrivalTime, _ := time.Parse(layout, arrival)
						departureTime, _ := time.Parse(layout, departure)

						difference := departureTime.Sub(arrivalTime)

						if arrivalTime.Before(departureTime) {
							if difference.Hours() < 10 {
								addReturnFlights = append(addReturnFlights, flight)
								addReturnFlights = append(addReturnFlights, flight2)
								returnFlights1S = append(returnFlights1S, addReturnFlights)
							}
						}
					}
				}
			}

			// 2 STOPS
			result3 := repo.db.Table("flights").
				Where("(deleted_at IS NULL) and "+
					"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
					"('' = ? or lower(date_of_departure) LIKE ?) and "+
					"(business_class_remaining_seats >= ?)",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%", "%"+strings.ToLower(flyingTo)+"%",
					returning, "%"+strings.ToLower(returning)+"%",
					passengerNumber,
				).
				Find(&returnFlights2SHelp)

			if result3.Error != nil {
				return nil, totalResults, result.Error
			}

			for _, flight := range returnFlights2SHelp {
				result3 := repo.db.Table("flights").
					Where("(deleted_at IS NULL) and "+
						"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
						"(business_class_remaining_seats >= ?)",
						flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%", "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
						passengerNumber,
					).
					Find(&returnFlights2SHelp2)

				if result3.Error != nil {
					return nil, totalResults, result.Error
				}

				for _, flight2 := range returnFlights2SHelp2 {
					if flight2.PlaceOfArrival == flyingTo {
						continue
					}

					var flightsHelp2 []*model.Flight

					result3 := repo.db.Table("flights").
						Where("(deleted_at IS NULL) and "+
							"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
							"('' = ? or lower(place_of_arrival) LIKE ? or lower(place_of_arrival) = ?) and "+
							"(business_class_remaining_seats >= ?)",
							flight2.PlaceOfArrival, "%"+strings.ToLower(flight2.PlaceOfArrival)+"%", "%"+strings.ToLower(flight2.PlaceOfArrival)+"%",
							flyingFrom, "%"+strings.ToLower(flyingFrom)+"%", "%"+strings.ToLower(flyingFrom)+"%",
							passengerNumber,
						).
						Find(&flightsHelp2)

					if result3.Error != nil {
						return nil, totalResults, result.Error
					}

					if len(flightsHelp2) != 0 {
						for _, flight3 := range flightsHelp2 {

							var addFlights []*model.Flight

							layout := "2006-01-02 15:04"
							flightCalc := flight.DateOfArrival + " " + flight.TimeOfArrival
							flight2Calc := flight2.DateOfDeparture + " " + flight2.TimeOfDeparture
							flight3Calc := flight3.DateOfDeparture + " " + flight3.TimeOfDeparture
							flightTime, _ := time.Parse(layout, flightCalc)
							flight2Time, _ := time.Parse(layout, flight2Calc)
							flight3Time, _ := time.Parse(layout, flight3Calc)

							if flightTime.Before(flight2Time) && flight2Time.Before(flight3Time) {
								addFlights = append(addFlights, flight)
								addFlights = append(addFlights, flight2)
								addFlights = append(addFlights, flight3)
								returnFlights2S = append(returnFlights2S, addFlights)
							}
						}
					}
				}
			}

		} else if travelClass == 3 {

			// DIRECT FLIGHTS
			result := repo.db.Table("flights").
				Where("(deleted_at IS NULL) and "+
					"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
					"('' = ? or lower(place_of_arrival) LIKE ? or lower(place_of_arrival) = ?) and "+
					"('' = ? or lower(date_of_departure) LIKE ?) and "+
					"(first_class_remaining_seats >= ?)",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%", "%"+strings.ToLower(flyingTo)+"%",
					flyingFrom, "%"+strings.ToLower(flyingFrom)+"%", "%"+strings.ToLower(flyingFrom)+"%",
					returning, "%"+strings.ToLower(returning)+"%",
					passengerNumber,
				).
				Find(&returnFlightsDirect)

			if result.Error != nil {
				return nil, returnTotalResults, result.Error
			}

			// 1 STOP
			result2 := repo.db.Table("flights").
				Where("(deleted_at IS NULL) and "+
					"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
					"('' = ? or lower(date_of_departure) LIKE ?) and "+
					"(first_class_remaining_seats >= ?)",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%", "%"+strings.ToLower(flyingTo)+"%",
					returning, "%"+strings.ToLower(returning)+"%",
					passengerNumber,
				).
				Find(&returnFlights1SHelp)

			if result2.Error != nil {
				return nil, returnTotalResults, result.Error
			}

			for _, flight := range returnFlights1SHelp {

				var returnFlightsHelp2 []*model.Flight

				result3 := repo.db.Table("flights").
					Where("(deleted_at IS NULL) and "+
						"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
						"('' = ? or lower(place_of_arrival) LIKE ? or lower(place_of_arrival) = ?) and "+
						"(first_class_remaining_seats >= ?)",
						flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%", "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
						flyingFrom, "%"+strings.ToLower(flyingFrom)+"%", "%"+strings.ToLower(flyingFrom)+"%",
						passengerNumber,
					).
					Find(&returnFlightsHelp2)

				if result3.Error != nil {
					return nil, returnTotalResults, result.Error
				}

				if len(returnFlightsHelp2) != 0 {
					for _, flight2 := range returnFlightsHelp2 {

						var addReturnFlights []*model.Flight

						layout := "2006-01-02 15:04"
						arrival := flight.DateOfArrival + " " + flight.TimeOfArrival
						departure := flight2.DateOfDeparture + " " + flight2.TimeOfDeparture
						arrivalTime, _ := time.Parse(layout, arrival)
						departureTime, _ := time.Parse(layout, departure)

						difference := departureTime.Sub(arrivalTime)

						if arrivalTime.Before(departureTime) {
							if difference.Hours() < 10 {
								addReturnFlights = append(addReturnFlights, flight)
								addReturnFlights = append(addReturnFlights, flight2)
								returnFlights1S = append(returnFlights1S, addReturnFlights)
							}
						}
					}
				}
			}

			// 2 STOPS
			result3 := repo.db.Table("flights").
				Where("(deleted_at IS NULL) and "+
					"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
					"('' = ? or lower(date_of_departure) LIKE ?) and "+
					"(first_class_remaining_seats >= ?)",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%", "%"+strings.ToLower(flyingTo)+"%",
					returning, "%"+strings.ToLower(returning)+"%",
					passengerNumber,
				).
				Find(&returnFlights2SHelp)

			if result3.Error != nil {
				return nil, totalResults, result.Error
			}

			for _, flight := range returnFlights2SHelp {
				result3 := repo.db.Table("flights").
					Where("(deleted_at IS NULL) and "+
						"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
						"(first_class_remaining_seats >= ?)",
						flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%", "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
						passengerNumber,
					).
					Find(&returnFlights2SHelp2)

				if result3.Error != nil {
					return nil, totalResults, result.Error
				}

				for _, flight2 := range returnFlights2SHelp2 {
					if flight2.PlaceOfArrival == flyingTo {
						continue
					}

					var flightsHelp2 []*model.Flight

					result3 := repo.db.Table("flights").
						Where("(deleted_at IS NULL) and "+
							"('' = ? or lower(place_of_departure) LIKE ? or lower(place_of_departure) = ?) and "+
							"('' = ? or lower(place_of_arrival) LIKE ? or lower(place_of_arrival) = ?) and "+
							"(first_class_remaining_seats >= ?)",
							flight2.PlaceOfArrival, "%"+strings.ToLower(flight2.PlaceOfArrival)+"%", "%"+strings.ToLower(flight2.PlaceOfArrival)+"%",
							flyingFrom, "%"+strings.ToLower(flyingFrom)+"%", "%"+strings.ToLower(flyingFrom)+"%",
							passengerNumber,
						).
						Find(&flightsHelp2)

					if result3.Error != nil {
						return nil, totalResults, result.Error
					}

					if len(flightsHelp2) != 0 {
						for _, flight3 := range flightsHelp2 {

							var addFlights []*model.Flight

							layout := "2006-01-02 15:04"
							flightCalc := flight.DateOfArrival + " " + flight.TimeOfArrival
							flight2Calc := flight2.DateOfDeparture + " " + flight2.TimeOfDeparture
							flight3Calc := flight3.DateOfDeparture + " " + flight3.TimeOfDeparture
							flightTime, _ := time.Parse(layout, flightCalc)
							flight2Time, _ := time.Parse(layout, flight2Calc)
							flight3Time, _ := time.Parse(layout, flight3Calc)

							if flightTime.Before(flight2Time) && flight2Time.Before(flight3Time) {
								addFlights = append(addFlights, flight)
								addFlights = append(addFlights, flight2)
								addFlights = append(addFlights, flight3)
								returnFlights2S = append(returnFlights2S, addFlights)
							}
						}
					}
				}
			}
		}

		for _, flight := range returnFlightsDirect {
			if flight.EconomyClassRemainingSeats == 0 && flight.BusinessClassRemainingSeats == 0 &&
				flight.FirstClassRemainingSeats == 0 {
				flight.FlightStatus = model.FULL
				repo.db.Table("flights").Save(&flight)
			}
		}

		for _, flight := range flightsDirect {

			for _, returnFlight := range returnFlightsDirect {

				layout := "2006-01-02 15:04"
				flightTo := flight.DateOfArrival + " " + flight.TimeOfArrival
				flightFrom := returnFlight.DateOfDeparture + " " + returnFlight.TimeOfDeparture
				flightToTime, _ := time.Parse(layout, flightTo)
				flightFromTime, _ := time.Parse(layout, flightFrom)

				if flightToTime.Before(flightFromTime) {
					var appendFlights [][]model.FlightDTO

					var flightList []model.FlightDTO
					flightList = append(flightList, flight.ToFlightDTO())

					var returnFlightList []model.FlightDTO
					returnFlightList = append(returnFlightList, returnFlight.ToFlightDTO())

					appendFlights = append(appendFlights, flightList)
					appendFlights = append(appendFlights, returnFlightList)

					allFlights = append(allFlights, appendFlights)
				}
			}
		}

		for _, flight := range flights1S {

			for _, returnFlight := range returnFlightsDirect {

				layout := "2006-01-02 15:04"
				flightTo := flight[1].DateOfArrival + " " + flight[1].TimeOfArrival
				flightFrom := returnFlight.DateOfDeparture + " " + returnFlight.TimeOfDeparture
				flightToTime, _ := time.Parse(layout, flightTo)
				flightFromTime, _ := time.Parse(layout, flightFrom)

				if flightToTime.Before(flightFromTime) {
					var appendFlights [][]model.FlightDTO

					var flightList []model.FlightDTO
					for _, flight2 := range flight {
						flightList = append(flightList, flight2.ToFlightDTO())
					}

					var returnFlightList []model.FlightDTO
					returnFlightList = append(returnFlightList, returnFlight.ToFlightDTO())

					appendFlights = append(appendFlights, flightList)
					appendFlights = append(appendFlights, returnFlightList)

					allFlights = append(allFlights, appendFlights)
				}
			}
		}

		for _, flight := range flightsDirect {

			for _, returnFlight := range returnFlights1S {

				layout := "2006-01-02 15:04"
				flightTo := flight.DateOfArrival + " " + flight.TimeOfArrival
				flightFrom := returnFlight[0].DateOfDeparture + " " + returnFlight[0].TimeOfDeparture
				flightToTime, _ := time.Parse(layout, flightTo)
				flightFromTime, _ := time.Parse(layout, flightFrom)

				if flightToTime.Before(flightFromTime) {

					var appendFlights [][]model.FlightDTO

					var flightList []model.FlightDTO
					flightList = append(flightList, flight.ToFlightDTO())

					var returnFlightList []model.FlightDTO
					for _, returnFlight2 := range returnFlight {
						returnFlightList = append(returnFlightList, returnFlight2.ToFlightDTO())
					}

					appendFlights = append(appendFlights, flightList)
					appendFlights = append(appendFlights, returnFlightList)

					allFlights = append(allFlights, appendFlights)
				}
			}
		}

		for _, flight := range flights1S {

			for _, returnFlight := range returnFlights1S {

				layout := "2006-01-02 15:04"
				flightTo := flight[1].DateOfArrival + " " + flight[1].TimeOfArrival
				flightFrom := returnFlight[0].DateOfDeparture + " " + returnFlight[0].TimeOfDeparture
				flightToTime, _ := time.Parse(layout, flightTo)
				flightFromTime, _ := time.Parse(layout, flightFrom)

				if flightToTime.Before(flightFromTime) {

					var appendFlights [][]model.FlightDTO

					var flightList []model.FlightDTO
					for _, flight2 := range flight {
						flightList = append(flightList, flight2.ToFlightDTO())
					}

					var returnFlightList []model.FlightDTO
					for _, returnFlight2 := range returnFlight {
						returnFlightList = append(returnFlightList, returnFlight2.ToFlightDTO())
					}

					appendFlights = append(appendFlights, flightList)
					appendFlights = append(appendFlights, returnFlightList)

					allFlights = append(allFlights, appendFlights)
				}
			}
		}

		for _, flight := range flights2S {

			for _, returnFlight := range returnFlightsDirect {

				layout := "2006-01-02 15:04"
				flightTo := flight[2].DateOfArrival + " " + flight[2].TimeOfArrival
				flightFrom := returnFlight.DateOfDeparture + " " + returnFlight.TimeOfDeparture
				flightToTime, _ := time.Parse(layout, flightTo)
				flightFromTime, _ := time.Parse(layout, flightFrom)

				if flightToTime.Before(flightFromTime) {

					var appendFlights [][]model.FlightDTO

					var flightList []model.FlightDTO
					for _, flight2 := range flight {
						flightList = append(flightList, flight2.ToFlightDTO())
					}

					var returnFlightList []model.FlightDTO
					returnFlightList = append(returnFlightList, returnFlight.ToFlightDTO())

					appendFlights = append(appendFlights, flightList)
					appendFlights = append(appendFlights, returnFlightList)

					allFlights = append(allFlights, appendFlights)
				}
			}
		}

		for _, flight := range flightsDirect {

			for _, returnFlight := range returnFlights2S {

				layout := "2006-01-02 15:04"
				flightTo := flight.DateOfArrival + " " + flight.TimeOfArrival
				flightFrom := returnFlight[0].DateOfDeparture + " " + returnFlight[0].TimeOfDeparture
				flightToTime, _ := time.Parse(layout, flightTo)
				flightFromTime, _ := time.Parse(layout, flightFrom)

				if flightToTime.Before(flightFromTime) {

					var appendFlights [][]model.FlightDTO

					var flightList []model.FlightDTO
					flightList = append(flightList, flight.ToFlightDTO())

					var returnFlightList []model.FlightDTO
					for _, returnFlight2 := range returnFlight {
						returnFlightList = append(returnFlightList, returnFlight2.ToFlightDTO())
					}

					appendFlights = append(appendFlights, flightList)
					appendFlights = append(appendFlights, returnFlightList)

					allFlights = append(allFlights, appendFlights)
				}
			}
		}

		for _, flight := range flights2S {

			for _, returnFlight := range returnFlights2S {

				layout := "2006-01-02 15:04"
				flightTo := flight[2].DateOfArrival + " " + flight[2].TimeOfArrival
				flightFrom := returnFlight[0].DateOfDeparture + " " + returnFlight[0].TimeOfDeparture
				flightToTime, _ := time.Parse(layout, flightTo)
				flightFromTime, _ := time.Parse(layout, flightFrom)

				if flightToTime.Before(flightFromTime) {

					var appendFlights [][]model.FlightDTO

					var flightList []model.FlightDTO
					for _, flight2 := range flight {
						flightList = append(flightList, flight2.ToFlightDTO())
					}

					var returnFlightList []model.FlightDTO
					for _, returnFlight2 := range returnFlight {
						returnFlightList = append(returnFlightList, returnFlight2.ToFlightDTO())
					}

					appendFlights = append(appendFlights, flightList)
					appendFlights = append(appendFlights, returnFlightList)

					allFlights = append(allFlights, appendFlights)
				}
			}
		}

		for _, flight := range flights1S {

			for _, returnFlight := range returnFlights2S {

				layout := "2006-01-02 15:04"
				flightTo := flight[1].DateOfArrival + " " + flight[1].TimeOfArrival
				flightFrom := returnFlight[0].DateOfDeparture + " " + returnFlight[0].TimeOfDeparture
				flightToTime, _ := time.Parse(layout, flightTo)
				flightFromTime, _ := time.Parse(layout, flightFrom)

				if flightToTime.Before(flightFromTime) {

					var appendFlights [][]model.FlightDTO

					var flightList []model.FlightDTO
					for _, flight2 := range flight {
						flightList = append(flightList, flight2.ToFlightDTO())
					}

					var returnFlightList []model.FlightDTO
					for _, returnFlight2 := range returnFlight {
						returnFlightList = append(returnFlightList, returnFlight2.ToFlightDTO())
					}

					appendFlights = append(appendFlights, flightList)
					appendFlights = append(appendFlights, returnFlightList)

					allFlights = append(allFlights, appendFlights)
				}
			}
		}

		for _, flight := range flights2S {

			for _, returnFlight := range returnFlights1S {

				layout := "2006-01-02 15:04"
				flightTo := flight[2].DateOfArrival + " " + flight[2].TimeOfArrival
				flightFrom := returnFlight[0].DateOfDeparture + " " + returnFlight[0].TimeOfDeparture
				flightToTime, _ := time.Parse(layout, flightTo)
				flightFromTime, _ := time.Parse(layout, flightFrom)

				if flightToTime.Before(flightFromTime) {

					var appendFlights [][]model.FlightDTO

					var flightList []model.FlightDTO
					for _, flight2 := range flight {
						flightList = append(flightList, flight2.ToFlightDTO())
					}

					var returnFlightList []model.FlightDTO
					for _, returnFlight2 := range returnFlight {
						returnFlightList = append(returnFlightList, returnFlight2.ToFlightDTO())
					}

					appendFlights = append(appendFlights, flightList)
					appendFlights = append(appendFlights, returnFlightList)

					allFlights = append(allFlights, appendFlights)
				}
			}
		}

		return allFlights, int64(len(allFlights)), nil
	}

	for _, flight := range flightsDirect {
		var flightList [][]model.FlightDTO
		var appendFlightList []model.FlightDTO
		var appendEmptyList []model.FlightDTO

		appendFlightList = append(appendFlightList, flight.ToFlightDTO())

		flightList = append(flightList, appendFlightList)
		flightList = append(flightList, appendEmptyList)

		allFlights = append(allFlights, flightList)
	}

	for _, flight := range flights1S {

		var flightList [][]model.FlightDTO
		var appendFlightList []model.FlightDTO
		var appendEmptyList []model.FlightDTO

		for _, flight2 := range flight {
			appendFlightList = append(appendFlightList, flight2.ToFlightDTO())
		}

		flightList = append(flightList, appendFlightList)
		flightList = append(flightList, appendEmptyList)

		allFlights = append(allFlights, flightList)
	}

	for _, flight := range flights2S {

		var flightList [][]model.FlightDTO
		var appendFlightList []model.FlightDTO
		var appendEmptyList []model.FlightDTO

		for _, flight2 := range flight {
			appendFlightList = append(appendFlightList, flight2.ToFlightDTO())
		}

		flightList = append(flightList, appendFlightList)
		flightList = append(flightList, appendEmptyList)

		allFlights = append(allFlights, flightList)
	}

	return allFlights, totalResults, nil
}

func dijkstra(graph *WeightedGraph, flyingFrom string) {

	visited := make(map[string]bool)
	heap := &Heap{}

	startNode := graph.GetNode(flyingFrom)
	startNode.value = 0
	heap.Push(startNode)

	for heap.Size() > 0 {
		current := heap.Pop()
		visited[current.name] = true
		edges := graph.Edges[current.name]
		for _, edge := range edges {
			if !visited[edge.node.name] {
				heap.Push(edge.node)
				if current.value+edge.weight < edge.node.value {
					edge.node.value = current.value + edge.weight
					edge.node.through = current
				}
			}
		}
	}
}

func buildGraph(cities []string, flights [][][]model.FlightDTO) *WeightedGraph {

	graph := NewGraph()
	nodes := AddNodes(graph, cities)

	for _, flight := range flights {
		for _, f := range flight[0] {
			graph.AddEdge(nodes[f.PlaceOfDeparture], nodes[f.PlaceOfArrival], int(f.EconomyClassPrice), f)
		}
	}

	return graph
}

// -- Weighted Graph

type Node struct {
	name    string
	value   int
	through *Node
}

type Edge struct {
	node   *Node
	weight int
	flight model.FlightDTO
}

type WeightedGraph struct {
	Nodes []*Node
	Edges map[string][]*Edge
	mutex sync.RWMutex
}

func NewGraph() *WeightedGraph {
	return &WeightedGraph{
		Edges: make(map[string][]*Edge),
	}
}

func (g *WeightedGraph) GetNode(name string) (node *Node) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	for _, n := range g.Nodes {
		if n.name == name {
			node = n
		}
	}
	return
}

func (g *WeightedGraph) AddNode(n *Node) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.Nodes = append(g.Nodes, n)
}

func AddNodes(graph *WeightedGraph, names []string) (nodes map[string]*Node) {
	nodes = make(map[string]*Node)
	for _, name := range names {
		n := &Node{name, math.MaxInt, nil}
		graph.AddNode(n)
		nodes[name] = n
	}
	return
}

func (g *WeightedGraph) AddEdge(n1, n2 *Node, weight int, flight model.FlightDTO) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	var allEdges = g.Edges[n1.name]

	if len(allEdges) == 0 {
		g.Edges[n1.name] = append(g.Edges[n1.name], &Edge{n2, weight, flight})

	} else {
		var contains = false
		for _, e := range allEdges {
			if e.flight.FlightNumber == flight.FlightNumber {
				contains = true
				break
			}
		}
		if !contains {
			g.Edges[n1.name] = append(g.Edges[n1.name], &Edge{n2, weight, flight})
		}
	}
	// g.Edges[n2.name] = append(g.Edges[n2.name], &Edge{n1, weight, flight})
}

func (n *Node) String() string {
	return n.name
}

func (e *Edge) String() string {
	return e.node.String() + "(" + strconv.Itoa(e.weight) + ")"
}

func (g *WeightedGraph) String() (s string) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	for _, n := range g.Nodes {
		s = s + n.String() + " ->"
		for _, c := range g.Edges[n.name] {
			s = s + " " + c.node.String() + " (" + strconv.Itoa(c.weight) + ")"
		}
		s = s + "\n"
	}
	return
}

// -- Heap

type Heap struct {
	elements []*Node
	mutex    sync.RWMutex
}

func (h *Heap) Size() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return len(h.elements)
}

// push an element to the heap, re-arrange the heap
func (h *Heap) Push(element *Node) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.elements = append(h.elements, element)
	i := len(h.elements) - 1
	for ; h.elements[i].value < h.elements[parent(i)].value; i = parent(i) {
		h.swap(i, parent(i))
	}
}

// pop the top of the heap, which is the min value
func (h *Heap) Pop() (i *Node) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	i = h.elements[0]
	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]
	h.rearrange(0)
	return
}

// rearrange the heap
func (h *Heap) rearrange(i int) {
	smallest := i
	left, right, size := leftChild(i), rightChild(i), len(h.elements)
	if left < size && h.elements[left].value < h.elements[smallest].value {
		smallest = left
	}
	if right < size && h.elements[right].value < h.elements[smallest].value {
		smallest = right
	}
	if smallest != i {
		h.swap(i, smallest)
		h.rearrange(smallest)
	}
}

func (h *Heap) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func (h *Heap) String() (str string) {
	return fmt.Sprintf("%q\n", getNames(h.elements))
}

func getNames(nodes []*Node) (names []string) {
	for _, node := range nodes {
		names = append(names, node.name)
	}
	return
}

func (repo *Repository) SortFlights(r *http.Request) ([]string, error) {

	var flights, _, _ = repo.FindAllByParams(r)
	queryParams := r.URL.Query()
	flyingFrom := queryParams.Get("flyingFrom")

	var allCities []string
	var numEdges = 0
	for _, flight := range flights {
		for _, f := range flight[0] {
			numEdges++
			allCities = append(allCities, f.PlaceOfDeparture)
			allCities = append(allCities, f.PlaceOfArrival)
		}
	}

	var cities []string
	for _, city := range allCities {
		var contains = false
		for _, c := range cities {
			if c == city {
				contains = true
			}
		}
		if !contains {
			cities = append(cities, city)
		}
	}

	graph := buildGraph(cities, flights)

	dijkstra(graph, flyingFrom)

	// display the nodes
	fmt.Println()

	var cheapestFlights []string
	for _, node := range graph.Nodes {
		if node.name != flyingFrom {
			var flightString = "Fly from " + flyingFrom + " to " + node.name + " for " + strconv.Itoa(node.value) + "$!"
			cheapestFlights = append(cheapestFlights, flightString)

			fmt.Printf("Cheapest flight from %s to %s is %d\n", flyingFrom, node.name, node.value)
			for n := node; n.through != nil; n = n.through {
				fmt.Print(n, " <- ")
			}
			fmt.Println(flyingFrom)
			fmt.Println()
		}
	}

	return cheapestFlights, nil
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
	flight.TimeOfBoarding = flightDTO.TimeOfBoarding
	flight.EconomyClassPoints = flightDTO.EconomyClassPoints
	flight.BusinessClassPoints = flightDTO.BusinessClassPoints
	flight.FirstClassPoints = flightDTO.FirstClassPoints

	result2 := repo.db.Table("flights").Save(&flight)

	if result2.Error != nil {
		return nil, errors.New("Flight cannot be updated!")
	}

	var retValue model.FlightDTO = flight.ToFlightDTO()
	return &retValue, nil
}
