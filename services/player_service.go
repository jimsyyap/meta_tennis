package services

import (
	"backend/models"
	"backend/repositories"
)

type PlayerService struct {
	playerRepository *repositories.PlayerRepository
}

func NewPlayerService(playerRepository *repositories.PlayerRepository) *PlayerService {
	return &PlayerService{playerRepository: playerRepository}
}

func (ps *PlayerService) GetPlayer(id int) (*models.Player, error) {
	return ps.playerRepository.GetPlayer(id)
}

func (ps *PlayerService) CreatePlayer(player *models.Player) error {
	return ps.playerRepository.CreatePlayer(player)
}
