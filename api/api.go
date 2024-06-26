package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// API is the API struct.
type API struct {
	DB *sql.DB
}

// NewAPI returns a new API struct.
func NewAPI(db *sql.DB) API {
	return API{
		DB: db,
	}
}

// ErrorResponse is the error response struct.
type ErrorResponse struct {
	Message string `json:"message"`
}

// sendErrorResponse sends an error response.
func sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}
