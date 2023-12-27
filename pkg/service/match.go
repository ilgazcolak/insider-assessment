package service

import (
	"github.com/ilgazcolak/insider-assessment/internal/repository"
	"github.com/ilgazcolak/insider-assessment/pkg/models"
)

type matchService struct {
	MatchRepository repository.MatchRepository
}

type MatchService interface {
	GetAll(map[int]string) (*[]models.Match, error)
}

func NewMatchService(matchRepository repository.MatchRepository) MatchService {
	return &matchService{MatchRepository: matchRepository}
}

func (s *matchService) GetAll(teamMap map[int]string) (*[]models.Match, error) {
	m, err := s.MatchRepository.GetAll()

	if err != nil {
		return nil, err
	}

	matches := make([]models.Match, 0)

	for _, mm := range m {
		var match models.Match

		match.Id = mm.Id
		match.HomeTeamId = mm.HomeTeamId
		match.HomeTeamName = teamMap[mm.HomeTeamId]
		match.AwayTeamId = mm.AwayTeamId
		match.AwayTeamName = teamMap[mm.AwayTeamId]

		matches = append(matches, match)
	}

	return &matches, nil
}
