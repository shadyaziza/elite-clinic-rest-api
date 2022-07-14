package server

import (
	"database/sql"
	db "github.com/shadyaziza/elite-clinic-rest-api/internal/db/sqlc"
)

type Store interface {
	db.Querier
}

type SQLStore struct {
	*db.Queries
	db *sql.DB
}

func NewStore(database *sql.DB) Store {
	return &SQLStore{
		db:      database,
		Queries: db.New(database),
	}
}
