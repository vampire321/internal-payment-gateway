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
	
	//verify connection to the database with a timeout context and with retry logic(docker safe)
	ctx, Cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer Cancel()
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
//configure connection pool settings for optimal performance
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := migrate(db); err != nil {
		return nil, err
	}
	return db, nil
}

	func migrate(db *sql.DB) error {

	query := `
	CREATE TABLE IF NOT EXISTS payments (
		id SERIAL PRIMARY KEY,
		user_id TEXT NOT NULL,
		amount NUMERIC NOT NULL,
		currency TEXT NOT NULL,
		status TEXT NOT NULL,
		idempotency_key TEXT UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	return nil
}


	