package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(payload)
}

func Error(w http.ResponseWriter, status int, message string) {

	type errResp struct {
		Error string `json:"error"`
	}

	JSON(w, status, errResp{
		Error: message,
	})
}