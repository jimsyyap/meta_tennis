package repositories

import (
	"database/sql"
	"log"

	"backend/models"
)

type PlayerRepository struct {
	db *sql.DB
}

func NewPlayerRepository(db *sql.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

func (pr *PlayerRepository) GetPlayer(id int) (*models.Player, error) {
	row := pr.db.QueryRow("SELECT * FROM players WHERE id = $1", id)
	player := &models.Player{}
	err := row.Scan(&player.ID, &player.Name, &player.Email, &player.Password, &player.Experience, &player.PlayingStyle)
	return player, err
}

func (pr *PlayerRepository) CreatePlayer(player *models.Player) error {
	_, err := pr.db.Exec("INSERT INTO players (name, email, password, experience, playing_style) VALUES ($1, $2, $3, $4, $5)", player.Name, player.Email, player.Password, player.Experience, player.PlayingStyle)
	return err
}
