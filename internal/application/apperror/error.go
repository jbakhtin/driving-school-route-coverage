package apperror

import (
	"encoding/json"
	"fmt"
)

var (
	SystemErrorCode       = "000"
	NotFoundCode          = "001"
	UserAlreadyExistsCode = "002"
	BadRequestParamsCode  = "003"
)

var (
	NotFound          = New(nil, "Not found", NotFoundCode, "", nil)
	UserAlreadyExists = New(nil, "User already exists", UserAlreadyExistsCode, "", nil)
	UserNotFound      = New(nil, "User doesn't exist with this login", NotFoundCode, "", nil)
)

type Errors map[string]string

type AppError struct {
	err              error
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             string `json:"code,omitempty"`
	Errors           Errors `json:"errors,omitempty"`
}

func New(err error, message string, code string, developerMessage string, errors Errors) *AppError {
	return &AppError{
		err,
		message,
		developerMessage,
		code,
		errors,
	}
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %v : %s", e.Message, e.err, e.Code)
}

func (e *AppError) Unwrap() error {
	return e.err
}

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}

	return marshal
}

func systemError(err error) *AppError {
	return New(err, "system error", SystemErrorCode, err.Error(), nil)
}
