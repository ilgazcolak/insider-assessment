package repository

import (
	"database/sql"

	"github.com/ilgazcolak/insider-assessment/internal/entities"
)

type playerRepository struct {
	Db *sql.DB
}

type PlayerRepository interface {
	GetAll() ([]entities.Player, error)
}

func NewPlayerRepository(db *sql.DB) PlayerRepository {
	return &playerRepository{Db: db}
}

func (r *playerRepository) GetAll() ([]entities.Player, error) {
	query := "select * from players"
	rows, err := r.Db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var players []entities.Player

	for rows.Next() {
		var p entities.Player

		err := rows.Scan(&p.Id, &p.TeamId, &p.Name, &p.Position)

		if err != nil {
			return nil, err
		}

		players = append(players, p)
	}

	return players, nil
}
