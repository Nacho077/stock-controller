package errors

import "fmt"

type InternalServerError struct {
	message string
}

func (err InternalServerError) Error() string {
	return err.message
}

func NewInternalServerError(message string, internalMessage string) InternalServerError {
	fmt.Println(fmt.Sprintf("Error: %s, internalMessage: %s", message, internalMessage))

	return InternalServerError{
		message: message,
	}
}
