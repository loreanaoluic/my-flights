package repository

import (
	"errors"
	"net/http"

	"github.com/my-flights/ReservationService/model"

	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (repo *Repository) FindTicketsByUserId(id uint, r *http.Request) ([]model.TicketDTO, int64, error) {
	var tickets []model.Ticket
	var ticketsDTO []model.TicketDTO
	var totalElements int64

	departing := time.Now().Format("2006-01-02")

	result := repo.db.Table("tickets").Where("(deleted_at IS NULL) and (user_id = ?) and "+
		"(date_of_departure > ?)",
		id, departing).
		Find(&tickets)
	repo.db.Table("tickets").Count(&totalElements)

	if result.Error != nil {
		return nil, totalElements, result.Error
	}

	for _, ticket := range tickets {
		ticketsDTO = append(ticketsDTO, ticket.ToTicketDTO())
	}

	return ticketsDTO, totalElements, nil
}

func (repo *Repository) FindHistoryByUserId(id uint, r *http.Request) ([]model.TicketDTO, int64, error) {
	var tickets []model.Ticket
	var ticketsDTO []model.TicketDTO
	var totalElements int64

	departing := time.Now().Format("2006-01-02")

	result := repo.db.Table("tickets").Where("(deleted_at IS NULL) and (user_id = ?) and "+
		"(date_of_departure < ?)",
		id, departing).
		Find(&tickets)
	repo.db.Table("tickets").Count(&totalElements)

	if result.Error != nil {
		return nil, totalElements, result.Error
	}

	for _, ticket := range tickets {
		ticketsDTO = append(ticketsDTO, ticket.ToTicketDTO())
	}

	return ticketsDTO, totalElements, nil
}

func (repo *Repository) CreateTicket(ticket model.Ticket) (model.Ticket, error) {

	createdTicket := repo.db.Create(&ticket)

	if createdTicket.Error != nil {
		return ticket, createdTicket.Error
	}

	return ticket, nil
}

func (repo *Repository) DeleteTicket(id uint) (*model.TicketDTO, error) {
	var ticket model.Ticket
	result := repo.db.Table("tickets").Where("id = ?", id).Clauses(clause.Returning{}).Delete(&ticket)

	if result.Error != nil {
		return nil, errors.New("Ticket cannot be deleted!")
	}

	var retValue model.TicketDTO = ticket.ToTicketDTO()
	return &retValue, nil
}
