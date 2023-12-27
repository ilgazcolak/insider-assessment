package repository

import (
	"database/sql"

	"github.com/ilgazcolak/insider-assessment/internal/entities"
)

type teamRepository struct {
	Db *sql.DB
}

type TeamRepository interface {
	GetAll() ([]entities.Team, error)
}

func NewTeamRepository(db *sql.DB) TeamRepository {
	return &teamRepository{Db: db}
}

func (r *teamRepository) GetAll() ([]entities.Team, error) {
	query := "select * from teams"
	rows, err := r.Db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var teams []entities.Team

	for rows.Next() {
		var t entities.Team

		err := rows.Scan(&t.Id, &t.Name)

		if err != nil {
			return nil, err
		}

		teams = append(teams, t)
	}

	return teams, nil
}
