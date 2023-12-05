package main

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/yourdb?sslmode=disable")
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://path/to/your/migrations",
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
