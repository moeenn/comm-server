package utils

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, payload any) {
	encoder := json.NewEncoder(w)
	w.WriteHeader(http.StatusBadRequest)
	encoder.Encode(payload)
}
