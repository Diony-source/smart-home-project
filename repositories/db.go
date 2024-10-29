package repositories

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDatabase() error {
	dsn := "postgres://myuser:mypassword@localhost:5432/smarthome"
	var err error

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return err
	}

	config.MaxConns = 10
	config.MinConns = 1
	config.MaxConnLifetime = time.Hour

	DB, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return err
	}

	log.Println("Connected to Postgresql database")
	return nil
}
