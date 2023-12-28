package errors

type InternalServerError struct {
	message string
}

func (err InternalServerError) Error() string {
	return err.message
}

func NewInternalServerError(message string) InternalServerError {
	return InternalServerError{
		message: message,
	}
}
