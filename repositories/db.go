package repositories

import (
    "context"
    "fmt"
    "os"
    "github.com/jackc/pgx/v4/pgxpool"
    "log"
)

var DB *pgxpool.Pool

func ConnectDatabase() error {
    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", dbUser, dbPassword, dbHost, dbName)

    config, err := pgxpool.ParseConfig(dsn)
    if err != nil {
        return err
    }

    DB, err = pgxpool.ConnectConfig(context.Background(), config)
    if err != nil {
        return err
    }

    log.Println("Connected to PostgreSQL database")
    return nil
}
