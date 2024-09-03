package services

import (
	"backend/models"
	"backend/repositories"
)

type CourtService struct {
	courtRepository *repositories.CourtRepository
}

func NewCourtService(courtRepository *repositories.CourtRepository) *CourtService {
	return &CourtService{courtRepository: courtRepository}
}

func (cs *CourtService) GetCourt(id int) (*models.Court, error) {
	return cs.courtRepository.GetCourt(id)
}

func (cs *CourtService) CreateCourt(court *models.Court) error {
	return cs.courtRepository.CreateCourt(court)
}
