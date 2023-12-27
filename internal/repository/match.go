package repository

import (
	"database/sql"

	"github.com/ilgazcolak/insider-assessment/internal/entities"
)

type matchRepository struct {
	Db *sql.DB
}

type MatchRepository interface {
	Get(id int) (entities.Match, error)
	GetAll() ([]entities.Match, error)
}

func NewMatchRepository(db *sql.DB) MatchRepository {
	return &matchRepository{Db: db}
}

func (r *matchRepository) Get(id int) (entities.Match, error) {
	query := "select * from matches where id = $1"
	row := r.Db.QueryRow(query, id)

	var match entities.Match

	err := row.Scan(&match.Id, &match.HomeTeamId, &match.AwayTeamId)

	if err != nil {
		return match, err
	}

	return match, nil
}

func (r *matchRepository) GetAll() ([]entities.Match, error) {
	query := "select * from matches"
	rows, err := r.Db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var matches []entities.Match

	for rows.Next() {
		var m entities.Match

		err := rows.Scan(&m.Id, &m.HomeTeamId, &m.AwayTeamId)

		if err != nil {
			return nil, err
		}

		matches = append(matches, m)
	}

	return matches, nil
}
