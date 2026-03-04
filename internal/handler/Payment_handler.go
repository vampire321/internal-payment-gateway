package handler
import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"STRIPE/internal/service"
)
//this represnt client input
type CreatePaymentRequest struct {
	UserID string
	Amount float64
	currency string
}

type PaymentHandler struct {
	Service *service.PaymentService
}

func NewPaymentHandler(s *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{Service : s}
}
func (h *PaymentHandler) CreatePayment(w http.ResponseWriter , r http.Request){
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed,"method_not_allowed","Only Post allowed")
		return
}
	idempotencyKey := r.Header.Get("idempotency-Key")
	//payment system requires this
	var req CreatePaymentRequest 
	if err := json.NewDecoder(r.Body)
}