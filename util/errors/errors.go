package errors

import (
	"fmt"
	"net/http"

	"gitlab.com/renodesper/gokit-microservices/util/error"
)

var (
	// UnexpectedPanic returned when there is an unexpected panic
	UnexpectedPanic = error.NewError(http.StatusInternalServerError, "ER9999", fmt.Errorf("Unexpected panic"))

	// StatusNotFound returned when client try to access undefined endpoint
	StatusNotFound = error.NewError(http.StatusNotFound, "ER9998", fmt.Errorf(http.StatusText(http.StatusNotFound)))

	// StatusBadRequest returned when ...
	StatusBadRequest = error.NewError(http.StatusBadRequest, "ER9997", fmt.Errorf(http.StatusText(http.StatusBadRequest)))
)

var (
	// InvalidCursor returned when the cursor for pagination is invalid
	InvalidCursor = error.NewError(http.StatusNotFound, "ER9901", fmt.Errorf("Cursor is invalid"))
)

var (
	// InvalidRequest returned when ...
	InvalidRequest = error.NewError(http.StatusBadRequest, "AU1001", fmt.Errorf("Request is not valid"))

	// UnparsableJSON returned when ...
	UnparsableJSON = error.NewError(http.StatusBadRequest, "AU1002", fmt.Errorf("Failed on parsing JSON"))

	// UnreadableResponse returned when ...
	UnreadableResponse = error.NewError(http.StatusBadRequest, "AU1003", fmt.Errorf("Failed to read response"))
)
