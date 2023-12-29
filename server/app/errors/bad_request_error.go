package errors

import "fmt"

type BadRequestError struct {
	message string
}

func (err BadRequestError) Error() string {
	return err.message
}

func NewBadRequestError(message string, internalMessage string) BadRequestError {
	fmt.Println(fmt.Sprintf("Error: %s, internalMessage: %s", message, internalMessage))

	return BadRequestError{
		message: message,
	}
}
