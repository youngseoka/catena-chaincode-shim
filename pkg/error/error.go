package error

import (
	"fmt"
)

type Error struct {
	Code      uint32 `json:"code"`
	Codespace string `json:"codespace"`
	Log       string `json:"log"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("Code: [%d], Codespace: [%s], Log: [%s]", e.Code, e.Codespace, e.Log)
}

func NewError(codespace string, code uint32, format string, args ...interface{}) *Error {
	log := fmt.Sprintf(format, args...)
	return &Error{
		Code:      code,
		Codespace: codespace,
		Log:       log,
	}
}

func ErrOk() *Error {
	return NewError("", Ok, "")
}

func ErrUnknown(err error) *Error {
	return NewError("", UnknownError, "Unknown error - %s", err)
}

func ErrUnknownMethod(err error) *Error {
	return NewError("", UnknownError, "Unknown error - %s", err)
}

func ErrNotFound(err error) *Error {
	return NewError("", UnknownError, "Unknown error - %s", err)
}

func ErrJsonMarshal(err error) *Error {
	return NewError("", UnknownError, "Unknown error - %s", err)
}

func ErrJsonUnmarshal(err error) *Error {
	return NewError("", UnknownError, "Unknown error - %s", err)
}

func ErrConvertMapToStruct(err error) *Error {
	return NewError("", UnknownError, "Unknown error - %s", err)
}
