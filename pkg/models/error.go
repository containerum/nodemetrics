package models

import "fmt"

type Error struct {
	Message string `json:"error"`
}

func (err *Error) Error() string {
	return err.Message
}

func ErrorF(msg string, args ...interface{}) *Error {
	return &Error{
		Message: fmt.Sprintf(msg, args...),
	}
}
