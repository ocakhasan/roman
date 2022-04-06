// package error defines the errors used in http responses

package error

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	// ErrEmptyInput shows that given input is emptyx
	ErrEmptyInput = errors.New("please provide an input")

	// ErrInvalidInput shows that given input is not an integer
	ErrInvalidInput = errors.New("please provide an integer")

	// ErrMinBiggerThanMax shows request body is invalid
	ErrMinBiggerThanMax = errors.New("min value cannot be larger than max")
)

// ErrorResponse represents error response body.
type ErrorResponse struct {
	Err        string `json:"error"`
	StatusCode int    `json:"StatusCode"`
}

// WriteError writes error to http response
func WriteError(w http.ResponseWriter, err error) {
	var statusCode int
	switch err {
	case ErrEmptyInput:
		statusCode = http.StatusBadRequest
	case ErrInvalidInput:
		statusCode = http.StatusBadRequest
	case ErrMinBiggerThanMax:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}
	errorResponse := ErrorResponse{Err: err.Error(), StatusCode: statusCode}
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(&errorResponse)
}
