package middleware
import "net/http"

const apiKey = "secret-payment-key"

func APIKeyAuth(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Key := r.Header.Get("Authorization")
			if Key != apiKey {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("unauthorized"))
				return
			}
			next.ServeHTTP(w,r)
})
}