package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"STRIPE/internal/config"
	"STRIPE/internal/database"
	"STRIPE/internal/handler"
	"STRIPE/internal/repository"
	"STRIPE/internal/service"
)
//load the confiuration
func main(){
	cfg, err := config.Load()
	if err != nil{
		log.Fatalf("failed to load config: %v", err)
	}
}