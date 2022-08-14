package db

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (db *Database) MigrateDB() error {

	fmt.Println("Migrating the database ... ")
	driver, err := postgres.WithInstance(db.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create the postgres driver: %w", err)

	}
	m, err := migrate.NewWithDatabaseInstance("file:migrations", "postgres", driver)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {

			return fmt.Errorf("could not run up the migrations: %w", err)
		}
	}
	fmt.Println("Successfully migrated the database")
	return nil
}
