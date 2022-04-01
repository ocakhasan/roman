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

	// ErrUnsupportedMethod shows that requested method is not allowed
	ErrUnsupportedMethod = errors.New("this method is not allowed in current url")

	// ErrInvalidBody shows request body is invalid
	ErrInvalidBody = errors.New("request body is invalid")

	// ErrInvalidDeliveryPoint shows that given delivery point is invalid.
	ErrInvalidDeliveryPoint = errors.New("invalid delivery point name")

	// ErrEmptyBarcode shows that given barcode name is empty
	ErrEmptyBarcode = errors.New("empty barcode. Please provide a barcode")
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
	case ErrUnsupportedMethod:
		statusCode = http.StatusMethodNotAllowed
	case ErrInvalidBody:
		statusCode = http.StatusBadRequest
	case ErrEmptyInput:
		statusCode = http.StatusBadRequest
	case ErrInvalidInput:
		statusCode = http.StatusBadRequest
	case ErrInvalidDeliveryPoint:
		statusCode = http.StatusBadRequest
	case ErrEmptyBarcode:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}
	errorResponse := ErrorResponse{Err: err.Error(), StatusCode: statusCode}
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(&errorResponse)
}