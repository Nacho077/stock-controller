package errors

import (
	"net/http"
)

func HandleError(err error) (int, string) {
	switch err.(type) {
	case InternalServerError:
		return http.StatusInternalServerError, err.Error()
	default:
		return http.StatusInternalServerError, err.Error()
	}
}
