package repository

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
				"('' = ? or lower(place_of_departure) LIKE ?) and "+
				"('' = ? or lower(place_of_arrival) LIKE ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(economy_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
				flyingTo, "%"+strings.ToLower(flyingTo)+"%",
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
				"('' = ? or lower(place_of_departure) LIKE ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(economy_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
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
					"('' = ? or lower(place_of_departure) LIKE ?) and "+
					"('' = ? or lower(place_of_arrival) LIKE ?) and "+
					"(economy_class_remaining_seats >= ?)",
					flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%",
					passengerNumber,
				).
				Find(&flightsHelp2)

			if result3.Error != nil {
				return nil, totalResults, result.Error
			}

			var addFlights []*model.Flight

			if len(flightsHelp2) != 0 {
				for _, flight2 := range flightsHelp2 {
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
				"('' = ? or lower(place_of_departure) LIKE ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(economy_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
				departing, "%"+strings.ToLower(departing)+"%",
				passengerNumber,
			).
			Find(&flights2SHelp)

		if result3.Error != nil {
			return nil, totalResults, result.Error
		}

		for _, flight := range flights2SHelp {
			result3 := repo.db.Table("flights").
				Where("(deleted_at IS NULL) and "+
					"('' = ? or lower(place_of_departure) LIKE ?) and "+
					"(economy_class_remaining_seats >= ?)",
					flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
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
						"('' = ? or lower(place_of_departure) LIKE ?) and "+
						"('' = ? or lower(place_of_arrival) LIKE ?) and "+
						"(economy_class_remaining_seats >= ?)",
						flight2.PlaceOfArrival, "%"+strings.ToLower(flight2.PlaceOfArrival)+"%",
						flyingTo, "%"+strings.ToLower(flyingTo)+"%",
						passengerNumber,
					).
					Find(&flightsHelp2)

				if result3.Error != nil {
					return nil, totalResults, result.Error
				}

				var addFlights []*model.Flight

				if len(flightsHelp2) != 0 {
					for _, flight3 := range flightsHelp2 {
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

	} else if travelClass == 2 {

		// DIRECT FLIGHTS
		result := repo.db.Table("flights").
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
			Find(&flightsDirect)

		if result.Error != nil {
			return nil, totalResults, result.Error
		}

		// 1 STOP
		result2 := repo.db.Table("flights").
			Where("(deleted_at IS NULL) and "+
				"('' = ? or lower(place_of_departure) LIKE ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(business_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
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
					"('' = ? or lower(place_of_departure) LIKE ?) and "+
					"('' = ? or lower(place_of_arrival) LIKE ?) and "+
					"(business_class_remaining_seats >= ?)",
					flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%",
					passengerNumber,
				).
				Find(&flightsHelp2)

			if result3.Error != nil {
				return nil, totalResults, result.Error
			}

			var addFlights []*model.Flight

			if len(flightsHelp2) != 0 {
				for _, flight2 := range flightsHelp2 {
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
				"('' = ? or lower(place_of_departure) LIKE ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(business_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
				departing, "%"+strings.ToLower(departing)+"%",
				passengerNumber,
			).
			Find(&flights2SHelp)

		if result3.Error != nil {
			return nil, totalResults, result.Error
		}

		for _, flight := range flights2SHelp {
			result3 := repo.db.Table("flights").
				Where("(deleted_at IS NULL) and "+
					"('' = ? or lower(place_of_departure) LIKE ?) and "+
					"(business_class_remaining_seats >= ?)",
					flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
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
						"('' = ? or lower(place_of_departure) LIKE ?) and "+
						"('' = ? or lower(place_of_arrival) LIKE ?) and "+
						"(business_class_remaining_seats >= ?)",
						flight2.PlaceOfArrival, "%"+strings.ToLower(flight2.PlaceOfArrival)+"%",
						flyingTo, "%"+strings.ToLower(flyingTo)+"%",
						passengerNumber,
					).
					Find(&flightsHelp2)

				if result3.Error != nil {
					return nil, totalResults, result.Error
				}

				var addFlights []*model.Flight

				if len(flightsHelp2) != 0 {
					for _, flight3 := range flightsHelp2 {
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
				"('' = ? or lower(place_of_departure) LIKE ?) and "+
				"('' = ? or lower(place_of_arrival) LIKE ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(first_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
				flyingTo, "%"+strings.ToLower(flyingTo)+"%",
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
				"('' = ? or lower(place_of_departure) LIKE ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(first_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
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
					"('' = ? or lower(place_of_departure) LIKE ?) and "+
					"('' = ? or lower(place_of_arrival) LIKE ?) and "+
					"(first_class_remaining_seats >= ?)",
					flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%",
					passengerNumber,
				).
				Find(&flightsHelp2)

			if result3.Error != nil {
				return nil, totalResults, result.Error
			}

			var addFlights []*model.Flight

			if len(flightsHelp2) != 0 {
				for _, flight2 := range flightsHelp2 {
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
				"('' = ? or lower(place_of_departure) LIKE ?) and "+
				"('' = ? or lower(date_of_departure) LIKE ?) and "+
				"(first_class_remaining_seats >= ?)",
				flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
				departing, "%"+strings.ToLower(departing)+"%",
				passengerNumber,
			).
			Find(&flights2SHelp)

		if result3.Error != nil {
			return nil, totalResults, result.Error
		}

		for _, flight := range flights2SHelp {
			result3 := repo.db.Table("flights").
				Where("(deleted_at IS NULL) and "+
					"('' = ? or lower(place_of_departure) LIKE ?) and "+
					"(first_class_remaining_seats >= ?)",
					flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
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
						"('' = ? or lower(place_of_departure) LIKE ?) and "+
						"('' = ? or lower(place_of_arrival) LIKE ?) and "+
						"(first_class_remaining_seats >= ?)",
						flight2.PlaceOfArrival, "%"+strings.ToLower(flight2.PlaceOfArrival)+"%",
						flyingTo, "%"+strings.ToLower(flyingTo)+"%",
						passengerNumber,
					).
					Find(&flightsHelp2)

				if result3.Error != nil {
					return nil, totalResults, result.Error
				}

				var addFlights []*model.Flight

				if len(flightsHelp2) != 0 {
					for _, flight3 := range flightsHelp2 {
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
					"('' = ? or lower(place_of_departure) LIKE ?) and "+
					"('' = ? or lower(place_of_arrival) LIKE ?) and "+
					"('' = ? or lower(date_of_departure) LIKE ?) and "+
					"(economy_class_remaining_seats >= ?)",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%",
					flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
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
					"('' = ? or lower(place_of_departure) LIKE ?) and "+
					"('' = ? or lower(date_of_departure) LIKE ?) and "+
					"(economy_class_remaining_seats >= ?)",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%",
					departing, "%"+strings.ToLower(departing)+"%",
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
						"('' = ? or lower(place_of_departure) LIKE ?) and "+
						"('' = ? or lower(place_of_arrival) LIKE ?) and "+
						"(economy_class_remaining_seats >= ?)",
						flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
						flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
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
					"('' = ? or lower(place_of_departure) LIKE ?) and "+
					"('' = ? or lower(date_of_departure) LIKE ?) and "+
					"(economy_class_remaining_seats >= ?)",
					flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
					departing, "%"+strings.ToLower(departing)+"%",
					passengerNumber,
				).
				Find(&returnFlights2SHelp)

			if result3.Error != nil {
				return nil, totalResults, result.Error
			}

			for _, flight := range returnFlights2SHelp {
				result3 := repo.db.Table("flights").
					Where("(deleted_at IS NULL) and "+
						"('' = ? or lower(place_of_departure) LIKE ?) and "+
						"(economy_class_remaining_seats >= ?)",
						flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
						passengerNumber,
					).
					Find(&returnFlights2SHelp2)

				if result3.Error != nil {
					return nil, totalResults, result.Error
				}

				for _, flight2 := range returnFlights2SHelp2 {
					var flightsHelp2 []*model.Flight

					result3 := repo.db.Table("flights").
						Where("(deleted_at IS NULL) and "+
							"('' = ? or lower(place_of_departure) LIKE ?) and "+
							"('' = ? or lower(place_of_arrival) LIKE ?) and "+
							"(economy_class_remaining_seats >= ?)",
							flight2.PlaceOfArrival, "%"+strings.ToLower(flight2.PlaceOfArrival)+"%",
							flyingTo, "%"+strings.ToLower(flyingTo)+"%",
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

							if flightTime.Before(flight2Time) && flight2Time.Before(flight3Time) && difference1.Hours() < 10 && difference2.Hours() < 10 {
								addFlights = append(addFlights, flight)
								addFlights = append(addFlights, flight2)
								addFlights = append(addFlights, flight3)
								returnFlights2S = append(returnFlights2S, addFlights)
							}
						}
					}
				}
			}

		} else if travelClass == 2 {

			// DIRECT FLIGHTS
			result := repo.db.Table("flights").
				Where("(deleted_at IS NULL) and "+
					"('' = ? or lower(place_of_departure) LIKE ?) and "+
					"('' = ? or lower(place_of_arrival) LIKE ?) and "+
					"('' = ? or lower(date_of_departure) LIKE ?) and "+
					"(business_class_remaining_seats >= ?)",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%",
					flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
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
					"('' = ? or lower(place_of_departure) LIKE ?) and "+
					"('' = ? or lower(date_of_departure) LIKE ?) and "+
					"(business_class_remaining_seats >= ?)",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%",
					departing, "%"+strings.ToLower(departing)+"%",
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
						"('' = ? or lower(place_of_departure) LIKE ?) and "+
						"('' = ? or lower(place_of_arrival) LIKE ?) and "+
						"(business_class_remaining_seats >= ?)",
						flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
						flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
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
					"('' = ? or lower(place_of_departure) LIKE ?) and "+
					"('' = ? or lower(date_of_departure) LIKE ?) and "+
					"(business_class_remaining_seats >= ?)",
					flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
					departing, "%"+strings.ToLower(departing)+"%",
					passengerNumber,
				).
				Find(&returnFlights2SHelp)

			if result3.Error != nil {
				return nil, totalResults, result.Error
			}

			for _, flight := range returnFlights2SHelp {
				result3 := repo.db.Table("flights").
					Where("(deleted_at IS NULL) and "+
						"('' = ? or lower(place_of_departure) LIKE ?) and "+
						"(business_class_remaining_seats >= ?)",
						flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
						passengerNumber,
					).
					Find(&returnFlights2SHelp2)

				if result3.Error != nil {
					return nil, totalResults, result.Error
				}

				for _, flight2 := range returnFlights2SHelp2 {
					var flightsHelp2 []*model.Flight

					result3 := repo.db.Table("flights").
						Where("(deleted_at IS NULL) and "+
							"('' = ? or lower(place_of_departure) LIKE ?) and "+
							"('' = ? or lower(place_of_arrival) LIKE ?) and "+
							"(business_class_remaining_seats >= ?)",
							flight2.PlaceOfArrival, "%"+strings.ToLower(flight2.PlaceOfArrival)+"%",
							flyingTo, "%"+strings.ToLower(flyingTo)+"%",
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
					"('' = ? or lower(place_of_departure) LIKE ?) and "+
					"('' = ? or lower(place_of_arrival) LIKE ?) and "+
					"('' = ? or lower(date_of_departure) LIKE ?) and "+
					"(first_class_remaining_seats >= ?)",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%",
					flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
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
					"('' = ? or lower(place_of_departure) LIKE ?) and "+
					"('' = ? or lower(date_of_departure) LIKE ?) and "+
					"(first_class_remaining_seats >= ?)",
					flyingTo, "%"+strings.ToLower(flyingTo)+"%",
					departing, "%"+strings.ToLower(departing)+"%",
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
						"('' = ? or lower(place_of_departure) LIKE ?) and "+
						"('' = ? or lower(place_of_arrival) LIKE ?) and "+
						"(first_class_remaining_seats >= ?)",
						flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
						flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
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
					"('' = ? or lower(place_of_departure) LIKE ?) and "+
					"('' = ? or lower(date_of_departure) LIKE ?) and "+
					"(first_class_remaining_seats >= ?)",
					flyingFrom, "%"+strings.ToLower(flyingFrom)+"%",
					departing, "%"+strings.ToLower(departing)+"%",
					passengerNumber,
				).
				Find(&returnFlights2SHelp)

			if result3.Error != nil {
				return nil, totalResults, result.Error
			}

			for _, flight := range returnFlights2SHelp {
				result3 := repo.db.Table("flights").
					Where("(deleted_at IS NULL) and "+
						"('' = ? or lower(place_of_departure) LIKE ?) and "+
						"(first_class_remaining_seats >= ?)",
						flight.PlaceOfArrival, "%"+strings.ToLower(flight.PlaceOfArrival)+"%",
						passengerNumber,
					).
					Find(&returnFlights2SHelp2)

				if result3.Error != nil {
					return nil, totalResults, result.Error
				}

				for _, flight2 := range returnFlights2SHelp2 {
					var flightsHelp2 []*model.Flight

					result3 := repo.db.Table("flights").
						Where("(deleted_at IS NULL) and "+
							"('' = ? or lower(place_of_departure) LIKE ?) and "+
							"('' = ? or lower(place_of_arrival) LIKE ?) and "+
							"(first_class_remaining_seats >= ?)",
							flight2.PlaceOfArrival, "%"+strings.ToLower(flight2.PlaceOfArrival)+"%",
							flyingTo, "%"+strings.ToLower(flyingTo)+"%",
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

	return repo.SortFlights(allFlights, totalResults)
}

type Graph struct {
	NumNodes int
	NumEdges int
	Nodes    []Node
	Edges    [][]Edge
}

type Node struct {
	Id   int
	City string
}

type Edge struct {
	From   int
	To     int
	Flight model.FlightDTO
	Weight int
}

func NewGraph(n, e int) *Graph {
	return &Graph{
		NumNodes: n,
		NumEdges: e,
		Nodes:    make([]Node, n),
		Edges:    make([][]Edge, e),
	}
}

func (g *Graph) AddNode(n Node) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(flight model.FlightDTO, from, to, w int) {
	g.Edges[from] = append(g.Edges[from], Edge{From: from, To: to, Flight: flight, Weight: w})
}

func (g *Graph) adjacentEdgesExample() {

	fmt.Printf("Printing all nodes (%d) in graph.\n", len(g.Nodes))
	for _, node := range g.Nodes {
		fmt.Printf("Node %d: %s\n", node.Id, node.City)
	}

	fmt.Println("\nPrinting all edges in graph.")
	for _, adjacent := range g.Edges {
		for _, e := range adjacent {
			fmt.Printf("Edge: %s -> %s (%d -> %d) (%d)\n", e.Flight.PlaceOfDeparture, e.Flight.PlaceOfArrival, e.From, e.To, e.Weight)
		}
	}
}

func (repo *Repository) SortFlights(flights [][][]model.FlightDTO, totalResults int64) ([][][]model.FlightDTO, int64, error) {

	// ONE WAY

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

	myGraph := NewGraph(0, numEdges)

	var nodeId = 0
	for _, city := range cities {
		myGraph.AddNode(Node{nodeId, city})
		nodeId++
	}

	for _, flight := range flights {
		for _, f := range flight[0] {
			var from = 0
			for _, city := range cities {
				if city == f.PlaceOfDeparture {
					break
				}
				from++
			}

			var to = 0
			for _, city := range cities {
				if city == f.PlaceOfArrival {
					break
				}
				to++
			}

			var containsEdge = false
			for _, edge := range myGraph.Edges {
				for _, e := range edge {
					if e.From == from && e.To == to && e.Weight == int(f.EconomyClassPrice) {
						containsEdge = true
					}
				}
			}
			if !containsEdge {
				myGraph.AddEdge(f, from, to, int(f.EconomyClassPrice))
				myGraph.NumEdges++
			}
		}
	}

	myGraph.adjacentEdgesExample()

	return flights, totalResults, nil
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
