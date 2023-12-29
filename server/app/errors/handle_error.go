package errors

import (
	"net/http"
)

func HandleError(err error) (int, string) {
	switch err.(type) {
	case FailedDependencyError:
		return http.StatusFailedDependency, err.Error()
	case BadRequestError:
		return http.StatusBadRequest, err.Error()
	default:
		return http.StatusInternalServerError, err.Error()
	}
}
