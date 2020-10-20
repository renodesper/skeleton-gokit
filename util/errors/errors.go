package errors

import (
	"fmt"
	"net/http"

	"gitlab.com/renodesper/gokit-microservices/util/error"
)

var (
	// UnexpectedPanic returned when there is an unexpected panic
	UnexpectedPanic = error.NewError(http.StatusInternalServerError, "AU9999", fmt.Errorf("Unexpected panic"))

	// StatusNotFound returned when client try to access undefined endpoint
	StatusNotFound = error.NewError(http.StatusNotFound, "AU9998", fmt.Errorf(http.StatusText(http.StatusNotFound)))

	// StatusBadRequest returned when ...
	StatusBadRequest = error.NewError(http.StatusBadRequest, "AU9997", fmt.Errorf(http.StatusText(http.StatusBadRequest)))
)
