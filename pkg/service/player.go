package service

import (
	"github.com/ilgazcolak/insider-assessment/internal/repository"
	"github.com/ilgazcolak/insider-assessment/pkg/models"
)

type playerService struct {
	PlayerRepository repository.PlayerRepository
}

type PlayerService interface {
	GetAll(map[int]string) (*[]models.Player, error)
}

func NewPlayerService(PlayerRepository repository.PlayerRepository) PlayerService {
	return &playerService{PlayerRepository: PlayerRepository}
}

func (s *playerService) GetAll(teamMap map[int]string) (*[]models.Player, error) {
	p, err := s.PlayerRepository.GetAll()

	if err != nil {
		return nil, err
	}

	players := make([]models.Player, 0)

	for _, pp := range p {
		var player models.Player

		player.Id = pp.Id
		player.TeamId = pp.TeamId
		player.TeamName = teamMap[pp.TeamId]
		player.Name = pp.Name
		player.Position = pp.Position

		players = append(players, player)
	}

	return &players, nil
}
