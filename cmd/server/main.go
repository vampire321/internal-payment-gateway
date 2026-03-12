package main

import (
	"context"
	"strconv"
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

	mux := http.NewServeMux()

	mux.HandleFunc("/payments", paymentHandler.CreatePayment)
		server := &http.Server{
		Addr:         ":" + strconv.Itoa(cfg.ServerPort),
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Printf("server running on port %d", cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Println("shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("shutdown error: %v", err)
	}
}