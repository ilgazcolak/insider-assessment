package service

import (
	"math"

	"github.com/ilgazcolak/insider-assessment/internal/repository"
	"github.com/ilgazcolak/insider-assessment/pkg/models"
)

type statisticsService struct {
	statisticRepository repository.StatisticsRepository
}

type StatisticsService interface {
	Get(playerId int, matchId int) (models.Statistics, error)
	GetAll(map[int]string) (*[]models.Statistics, error)
	Update(models.Statistics) error
}

func NewStatisticsService(statisticRepository repository.StatisticsRepository) StatisticsService {
	return &statisticsService{statisticRepository: statisticRepository}
}

func (s *statisticsService) Get(playerId int, matchId int) (models.Statistics, error) {
	ss, err := s.statisticRepository.Get(playerId, matchId)

	if err != nil {
		return models.Statistics{}, err
	}

	var statistics models.Statistics

	statistics.Id = ss.Id
	statistics.PlayerId = ss.PlayerId
	statistics.MatchId = ss.MatchId
	statistics.Score = ss.Score
	statistics.AttemptCount = ss.AttemptCount
	statistics.TwoPointsAttempt = ss.TwoPointsAttempt
	statistics.ThreePointsAttempt = ss.ThreePointsAttempt
	statistics.TwoPointsSuccess = ss.TwoPointsSuccess
	statistics.ThreePointsSuccess = ss.ThreePointsSuccess
	if ss.TwoPointsAttempt > 0 {
		statistics.TwoPointsRate = math.Floor((float64(ss.TwoPointsSuccess)/float64(ss.TwoPointsAttempt))*100) / 100
	}
	if ss.ThreePointsAttempt > 0 {
		statistics.ThreePointsRate = math.Floor((float64(ss.ThreePointsSuccess)/float64(ss.ThreePointsAttempt))*100) / 100
	}
	statistics.Assist = ss.Assist

	return statistics, nil
}

func (s *statisticsService) GetAll(playerMap map[int]string) (*[]models.Statistics, error) {
	sSlice, err := s.statisticRepository.GetAll()

	if err != nil {
		return nil, err
	}

	statistics := make([]models.Statistics, 0)

	for _, ss := range sSlice {
		var statistic models.Statistics

		statistic.Id = ss.Id
		statistic.PlayerId = ss.PlayerId
		statistic.PlayerName = playerMap[ss.PlayerId]
		statistic.MatchId = ss.MatchId
		statistic.Score = ss.Score
		statistic.AttemptCount = ss.AttemptCount
		statistic.TwoPointsAttempt = ss.TwoPointsAttempt
		statistic.ThreePointsAttempt = ss.ThreePointsAttempt
		statistic.TwoPointsSuccess = ss.TwoPointsSuccess
		statistic.ThreePointsSuccess = ss.ThreePointsSuccess
		if ss.TwoPointsAttempt > 0 {
			statistic.TwoPointsRate = math.Floor((float64(ss.TwoPointsSuccess)/float64(ss.TwoPointsAttempt))*100) / 100
		}
		if ss.ThreePointsAttempt > 0 {
			statistic.ThreePointsRate = math.Floor((float64(ss.ThreePointsSuccess)/float64(ss.ThreePointsAttempt))*100) / 100
		}
		statistic.Assist = ss.Assist

		statistics = append(statistics, statistic)
	}

	return &statistics, nil
}

func (s *statisticsService) Update(model models.Statistics) error {
	return s.statisticRepository.Update(model)
}
