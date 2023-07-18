package helpers

import (
	"encoding/json"
	"net/http"
)

// Response is a helper function to create a response
func Response(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// ErrorResponse is a helper function to create an error response
func ErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	Response(w, statusCode, map[string]string{
		"error": err.Error(),
	})
}
