package apperror

import (
	"encoding/json"
	"fmt"
)

var (
	NotFound          = New(nil, "Not found", "001", "")
	UserAlreadyExists = New(nil, "User already exists", "002", "")
)

type AppError struct {
	err              error
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             string `json:"code,omitempty"`
}

func New(err error, message string, code string, developerMessage string) *AppError {
	return &AppError{
		err,
		message,
		developerMessage,
		code,
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
	return New(err, "system error", "000", err.Error())
}
