package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	_ "github.com/lib/pq"
	"STRIPE/internal/config"
)

func Connect(cfg *config.Config) (*sql.DB, error){
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	//open connection pool to the database
	db, err := sql.Open("postgres", connStr)
		if err != nil {
			return nil, fmt.Errorf("failed to open database connection: %w", err)
		}
}