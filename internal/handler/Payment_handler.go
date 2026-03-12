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
func (h *PaymentHandler) CreatePayment(w http.ResponseWriter , r *http.Request){
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed,"method_not_allowed","Only Post allowed")
		return
}
	idempotencyKey := r.Header.Get("idempotency-Key")
	//payment system requires this
	var req CreatePaymentRequest 
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
		writeError(w,http.StatusBadRequest,"invalid_json","malformed JSON body")
		return
	}
	ctx,cancel := context.WithTimeout(r.Context(),5*time.Second) //if stripe or database takes more than 5 seconds to finish it is cutting the connection
	defer cancel()

	//official hands off responsibility to the Service
	Payment,err := h.Service.ProcessPayment(
		ctx,
		req.UserID,
		req.Amount,
		req.currency,
		idempotencyKey,
	)
	if err != nil{
		writeError(w,http.StatusUnprocessableEntity,"processing_err",err.Error())
		return
	}
	//simply passes the raw data to the process Payment
	w.Header().Set("content-Tpe","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Payment)
	}
	func writeError(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": map[string]string{
			"code":    code,
			"message": message,
		},
	})
}