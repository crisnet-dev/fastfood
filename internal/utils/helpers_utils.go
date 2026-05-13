package utils

import (
	"encoding/json"
	"net/http"
)

func HttpError(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(map[string]any{
		"error": message,
	})
}

func HttpResponse(w http.ResponseWriter, message any, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}
