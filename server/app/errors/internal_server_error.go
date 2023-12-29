package errors

import "fmt"

type InternalServerError struct {
	message string
}

func (err InternalServerError) Error() string {
	return err.message
}

func NewInternalServerError(message string, internalMessage string) InternalServerError {
	fmt.Println("Error: " + message + "internalMessage: " + internalMessage)

	return InternalServerError{
		message: message,
	}
}
