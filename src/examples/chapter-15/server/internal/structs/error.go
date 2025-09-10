package structs

import (
	"fmt"
	"time"
)

// ErrorResponse represents a structured error response
type ErrorResponse struct {
	ErrorDetail ErrorDetail `json:"error"`
	RequestID   string      `json:"requestId,omitempty"`
	Timestamp   string      `json:"timestamp"`
}

// ErrorDetail contains detailed error information
type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// NewErrorResponse creates a new error response
func NewErrorResponse(code, message, details, requestID string) *ErrorResponse {
	return &ErrorResponse{
		ErrorDetail: ErrorDetail{
			Code:    code,
			Message: message,
			Details: details,
		},
		RequestID: requestID,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
}

// Common error codes
const (
	ErrCodeInvalidRequest   = "INVALID_REQUEST"
	ErrCodeAuthentication   = "AUTHENTICATION_ERROR"
	ErrCodeAuthorization    = "AUTHORIZATION_ERROR"
	ErrCodeNotFound         = "NOT_FOUND"
	ErrCodeRateLimit        = "RATE_LIMIT_EXCEEDED"
	ErrCodeInternal         = "INTERNAL_ERROR"
	ErrCodeFlowExecution    = "FLOW_EXECUTION_ERROR"
	ErrCodeToolExecution    = "TOOL_EXECUTION_ERROR"
	ErrCodeModelUnavailable = "MODEL_UNAVAILABLE"
)

// Error implements the error interface
func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorDetail.Code, e.ErrorDetail.Message)
}
