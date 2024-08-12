package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	Client sqlx.DB
}

func NewDatabaseConnection() (*Database, error) {
	connectString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)
	dbConn, err := sqlx.Connect("postgres", connectString)
	if err != nil {
		return &Database{}, fmt.Errorf("couldn't connect to the database")
	}
	return &Database{Client: *dbConn}, nil
}

func (d *Database) PingDatabase(ctx context.Context) error {

	return d.Client.PingContext(ctx)
}
