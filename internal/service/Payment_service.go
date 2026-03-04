package service

import(
	"errors"
	"context"
	"STRIPE/internal/model"
	"STRIPE/internal/repository"
)

type PaymentService struct {
	repo *repository.PaymentRepository
}
func NewPaymentService(repo *repository.PaymentRepository) *PaymentService {
	return &PaymentService{repo : repo}
}
func (s *PaymentService) ProcessPayment(ctx context.Context ,userID string,amount float64, currency string, idempotency string,)(*model.Payment, error){
	if userID == ""{
		return nil, errors.New("user id is required")
	}
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
}
	if currency == ""{
		return nil, errors.New("currency is required")
	}
	if idempotency == ""{
		return nil, errors.New("idempotency key is required")
	}
	existingPayment, err := s.repo.GetByIdempotencyKey(ctx,idempotency)
	if err != nil{
		return nil,err
	}
	if existingPayment != nil{
		return existingPayment,nil
	}

	payment := model.NewPayment(
		userID,
		amount,
		currency,
		idempotency,
	)
	payment.Status = model.StatusProcessed
	err = s.repo.Create(ctx,payment) //payment is dtruct that handed over to the create method
	if err != nil {
		return nil,err
	}
	return payment,nil
}