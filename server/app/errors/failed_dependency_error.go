package errors

import "fmt"

type FailedDependencyError struct {
	message string
}

func (err FailedDependencyError) Error() string {
	return err.message
}

func NewFailedDependencyError(message string, internalMessage string) FailedDependencyError {
	fmt.Println("Error: ", message, "internalMessage: ", internalMessage)

	return FailedDependencyError{
		message: message,
	}
}
