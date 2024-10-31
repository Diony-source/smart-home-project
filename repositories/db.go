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

    // DSN formatı yalnızca bir "postgres://" içerecek şekilde ayarlanıyor
    dsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", dbUser, dbPassword, dbHost, dbName)
    
    // DSN çıktısını tekrar loglayarak kontrol edelim
    log.Println("Connecting to database with DSN:", dsn)

    config, err := pgxpool.ParseConfig(dsn)
    if err != nil {
        log.Println("Error parsing DSN:", err)
        return err
    }

    DB, err = pgxpool.ConnectConfig(context.Background(), config)
    if err != nil {
        log.Println("Error connecting to database:", err)
        return err
    }

    log.Println("Connected to PostgreSQL database")
    return nil
}
