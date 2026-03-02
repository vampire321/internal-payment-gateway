package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"STRIPE/internal/model"
)
//paymentRepository its only job is to handle storage of payments
type PaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}
func (r *PaymentRepository) Create(ctx context.Context, p *model.Payment) error {
	query := `INSERT INTO payments 
	         (user_id, amount, currency, status, idempotency_key, created_at)
			 VALUES($1,$2,$3,$4,$5,$6) 
			 RETURNING id , created_at;
			 `
	//ON CONFLICT (idempotency_key) 
	//DO UPDATE SET status = EXCLUDED.status
	err := r.db.QueryRowContext(
		ctx , 
		query, 
		p.UserID, p.Amount, p.Currency, p.Status, p.IdempotencyKey, p.CreatedAt,//The values p.UserID, p.Amount, p.Currency, p.Status, p.IdempotencyKey, p.CreatedAt are bound to the placeholders $1, $2, $3, $4, $5, $6 in the SQL query.


	).Scan(&p.ID, &p.CreatedAt)
	if err != nil {
		if isUniqueViolation(err){
			return fmt.Errorf("payment with idempotency key already exists: %w", err)
		}
		return err
	}
	return nil
}

func (r *PaymentRepository) GetByIdempotencyKey(ctx context.Context, key string) (*model.Payment, error) {
	query := `SELECT id, user_id, amount, currency, status, idempotency_key, created_at
			  FROM payments
			  WHERE idempotency_key = $1;`

    row := r.db.QueryRowContext(ctx,query,key)
	var p model.Payment
	err := row.Scan(
		&p.ID,
		&p.UserID,
		&p.Amount,
		&p.Currency,
		&p.Status,
		&p.IdempotencyKey,
		&p.CreatedAt,
	)
	if err != nil{
		if err == sql.ErrNoRows {
			return nil, nil // No payment found with the given idempotency key
		}
		return nil, err // something went wrong with the query or like db connection
	}
	return &p, nil // Payment found, return it
}
func isUniqueViolation(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "duplicate key")
}