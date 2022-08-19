package repository

import (
	"net/http"
	"strconv"

	"github.com/my-flights/ReservationService/model"

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

func (repo *Repository) FindTicketsByUserId(id uint, r *http.Request) ([]model.TicketDTO, int64, error) {
	var tickets []model.Ticket
	var ticketsDTO []model.TicketDTO
	var totalElements int64

	result := repo.db.Scopes(Paginate(r)).Table("tickets").Where("user_id = ?", id).Find(&tickets)
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
