package model

import "time"

type PaymentStatus string

//allowed status constants for payment processing

const (
	StatusPending PaymentStatus = "pending"
	StatusProcessed PaymentStatus = "processed"
	StatusFailed PaymentStatus = "failed"
)

type Payment struct {
	ID int `json:"id"`
	UserID string `json:"user_id"`
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
	Status PaymentStatus `json:"status"`
	IdempotencyKey string `json:"idempotency_key"`
	CreatedAt time.Time `json:"created_at"`
}

func NewPayment(userID string,amount float64,currency string,idempotencykey string) *Payment {
	return &Payment{
		UserID: userID,
		Amount: amount,
		Currency: currency,
		Status: StatusPending,
		IdempotencyKey: idempotencykey,
}
}