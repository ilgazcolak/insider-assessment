package service

import (
	"github.com/ilgazcolak/insider-assessment/internal/repository"
	"github.com/ilgazcolak/insider-assessment/pkg/models"
)

type teamService struct {
	teamRepository repository.TeamRepository
}

type TeamService interface {
	GetAll() (*[]models.Team, error)
}

func NewTeamService(teamRepository repository.TeamRepository) TeamService {
	return &teamService{teamRepository: teamRepository}
}

func (s *teamService) GetAll() (*[]models.Team, error) {
	t, err := s.teamRepository.GetAll()

	if err != nil {
		return nil, err
	}

	teams := make([]models.Team, 0)

	for _, tt := range t {
		var team models.Team

		team.Id = tt.Id
		team.Name = tt.Name

		teams = append(teams, team)
	}

	return &teams, nil
}
