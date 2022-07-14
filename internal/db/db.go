package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type Database struct {
	Client *sql.DB
}

func NewDatabase() (*Database, error) {
	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",

		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)
	pool, err := sql.Open("postgres", connString)
	if err != nil {
		return &Database{}, fmt.Errorf("could not connect to the database: %w", err)
	}

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)
	return &Database{
		Client: pool,
	}, nil
}

func (db *Database) Ping(ctx context.Context) error {
	return db.Client.PingContext(ctx)
}
