package internal

import (
	"database/sql"
	db "github.com/shadyaziza/elite-clinic-rest-api/internal/db/sqlc"
)

// Store - this interface defines all the methods
// that  our services needs to operate, all the methods
// will be compatible with db.Querier interface
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
