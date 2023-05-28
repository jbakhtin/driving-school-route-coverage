package middleware

type ErrorResponse struct {
	Errors map[string]string `json:"errors,omitempty"`
}

func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{
		Errors: make(map[string]string),
	}
}
