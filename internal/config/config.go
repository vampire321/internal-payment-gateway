package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DBHost string
	DBPort int
	DBUser string
	DBPassword string
	DBName string
	ServerPort int
}
//load and validate configuration from environment variables
func load() (*Config , error){
	dbportstr := getEnv("DB_PORT","5432")
	dbPort ,err := strconv.Atoi(dbportstr)
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %w",err)
	}
	serverPortstr := getEnv("PORT","8080")
	serverPort, err:= strconv.Atoi(serverPortstr)
	if err != nil{
		return nil, fmt.Errorf("Invalid Server Port: %w",err)
	}
cfg := &Config{
	DBHost: getEnv("DB_HOST","localhost"),
	DBPort: dbPort,
	DBUser: getEnv("DB_User","postgres"),
	DBPassword: getEnv("DB_PASSWORD","password"),
	DBName: getEnv("DB_NAME","payments"),
	ServerPort: serverPort,
}
	if cfg.DBHost == "" || cfg.DBUser == "" || cfg.DBName == "" {
		return nil, fmt.Errorf("missing required database configuration")
	}
	return cfg, nil
}

//HELPER FUNCTION TO GET ENVIRONMENT VARIABLES WITH DEFAULT VALUES
func getEnv(key, fallback string) string {
	if value,exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
