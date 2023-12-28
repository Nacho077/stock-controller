package errors

type FailedDependency struct {
	message string
}

func (err FailedDependency) Error() string {
	return err.message
}

func NewFailedDependency(message string) FailedDependency {
	return FailedDependency{
		message: message,
	}
}
