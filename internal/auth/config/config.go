package config

import (
    "database/sql"
    "fmt"
    "os"

    _ "github.com/lib/pq" // PostgreSQL driver
)

func GetDBConnection() (*sql.DB, error) {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)

    return sql.Open("postgres", dsn)
}