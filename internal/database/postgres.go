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
	//verify connection to the database with a timeout context and with retry logic
	ctx, Cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	for{
	err = db.PingContext(ctx)
	if err == nil{
		break
	}
	select {
	case <-ctx.Done():
		return nil , fmt.Errorf("database connection not reachable: %w", err)
	case <- time.After(2*time.Second):
	}
}