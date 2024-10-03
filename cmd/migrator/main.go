package main

import (
	"database/sql"
	"errors"
	"fmt"
	"polygames/internal/infrastructure/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println(config.Config.DatabaseConnString)
	db, err := sql.Open("postgres", config.Config.DatabaseConnString)
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver,
	)
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil && !errors.Is(migrate.ErrNoChange, err) {
		panic(err)
	}
}
