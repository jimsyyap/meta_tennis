package repositories

import (
	"database/sql"
	"log"

	"backend/models"
)

type CourtRepository struct {
	db *sql.DB
}

func NewCourtRepository(db *sql.DB) *CourtRepository {
	return &CourtRepository{db: db}
}

func (cr *CourtRepository) GetCourt(id int) (*models.Court, error) {
	row := cr.db.QueryRow("SELECT * FROM courts WHERE id = $1", id)
	court := &models.Court{}
	err := row.Scan(&court.ID, &court.Name, &court.Location, &court.Availability, &court.Price)
	return court, err
}

func (cr *CourtRepository) CreateCourt(court *models.Court) error {
	_, err := cr.db.Exec("INSERT INTO courts (name, location, availability, price) VALUES ($1, $2, $3, $4)", court.Name, court.Location, court.Availability, court.Price)
	return err
}
