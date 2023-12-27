package repository

import (
	"database/sql"

	"github.com/ilgazcolak/insider-assessment/internal/entities"
	"github.com/ilgazcolak/insider-assessment/pkg/models"
)

type statisticsRepository struct {
	Db *sql.DB
}

type StatisticsRepository interface {
	Get(playerId int, matchId int) (entities.Statistics, error)
	GetAll() ([]entities.Statistics, error)
	Update(models.Statistics) error
}

func NewStatisticsRepository(db *sql.DB) StatisticsRepository {
	return &statisticsRepository{Db: db}
}

func (r *statisticsRepository) Get(playerId int, matchId int) (entities.Statistics, error) {
	var s entities.Statistics
	query := "select s.*, t.name, p.position from statistics s join players p on s.player_id = p.id join teams t on t.id = p.team_id where player_id=$1 and match_id=$2"

	row := r.Db.QueryRow(query, playerId, matchId)

	err := row.Scan(&s.Id, &s.PlayerId, &s.MatchId, &s.Score, &s.AttemptCount, &s.TwoPointsAttempt, &s.ThreePointsAttempt, &s.TwoPointsSuccess, &s.ThreePointsSuccess, &s.Assist, &s.TeamName, &s.Position)

	return s, err
}

func (r *statisticsRepository) GetAll() ([]entities.Statistics, error) {
	query := "select s.*, t.name, p.position from statistics s join players p on s.player_id = p.id join teams t on t.id = p.team_id"

	rows, err := r.Db.Query(query)

	if err != nil {
		return nil, err
	}

	statistics := []entities.Statistics{}

	for rows.Next() {
		var s entities.Statistics

		err := rows.Scan(&s.Id, &s.PlayerId, &s.MatchId, &s.Score, &s.AttemptCount, &s.TwoPointsAttempt, &s.ThreePointsAttempt, &s.TwoPointsSuccess, &s.ThreePointsSuccess, &s.Assist, &s.TeamName, &s.Position)

		if err != nil {
			return nil, err
		}

		statistics = append(statistics, s)
	}

	return statistics, nil
}

func (r *statisticsRepository) Update(m models.Statistics) error {
	query := "update statistics set score=$3, attempt_count=$4, two_points_attempt=$5, three_points_attempt=$6, two_points_success=$7, three_points_success=$8, assist=$9 where player_id=$1 and match_id=$2"

	_, err := r.Db.Exec(query, m.PlayerId, m.MatchId, m.Score, m.AttemptCount, m.TwoPointsAttempt, m.ThreePointsAttempt, m.TwoPointsSuccess, m.ThreePointsSuccess, m.Assist)

	if err != nil {
		return err
	}

	return nil
}
