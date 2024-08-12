package db

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (d *Database) Migrate() {
	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error creating Postgres driver: %v", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres",
		driver)
	if err != nil {
		log.Fatalf("Error creating migration instance: %v", err)
		return
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error applying migrations: %v", err)
		return
	}

	fmt.Println("Migrations applied successfully")
}
