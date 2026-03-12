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
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	paymentRepo := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepo)
	paymentHandler := handler.NewPaymentHandler(paymentService)
}
	mux := http.NewServeMux()

	mux.HandleFunc("/payments", paymentHandler.CreatePayment)