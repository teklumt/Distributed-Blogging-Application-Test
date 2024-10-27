package domain

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type ErrorResponseDomain interface {
	NewErrorResponse(message string) ErrorResponse
}