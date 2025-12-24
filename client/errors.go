package client

import (
	"fmt"
	"net/http"
)

// APIError represents an error returned by the VersaFleet API
type APIError struct {
	StatusCode int         `json:"-"`
	Message    string      `json:"message"`
	Errors     interface{} `json:"errors,omitempty"` // Can be a map or list details
	RequestID  string      `json:"request_id,omitempty"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("versafleet-sdk: status=%d message=%s request_id=%s", e.StatusCode, e.Message, e.RequestID)
}

// IsNotFound checks if the error is a 404
func (e *APIError) IsNotFound() bool {
	return e.StatusCode == http.StatusNotFound
}

// IsRateLimited checks if the error is a 429
func (e *APIError) IsRateLimited() bool {
	return e.StatusCode == http.StatusTooManyRequests
}
