package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Db *sql.DB
}

func NewPostgres(dsn string) *Postgres {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("unable to open database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	return &Postgres{Db: db}
}
