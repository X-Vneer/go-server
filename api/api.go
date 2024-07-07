package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

// Coin balance response
type CoinBalanceResponse struct {
	// success code
	Code int

	// account balance
	Balance int64
}

// Error response
type Error struct {
	Code    int
	Message string
}

// error response func
func writeError(w http.ResponseWriter, message string, code int) {
	response := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)
