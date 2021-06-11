package errors

import (
	"fmt"
	"net/http"

	"gitlab.com/renodesper/gokit-microservices/util/error"
)

var (
	UnexpectedPanic  = error.NewError(http.StatusInternalServerError, "ER9999", fmt.Errorf("unexpected panic"))
	StatusNotFound   = error.NewError(http.StatusNotFound, "ER9998", fmt.Errorf(http.StatusText(http.StatusNotFound)))
	StatusBadRequest = error.NewError(http.StatusBadRequest, "ER9997", fmt.Errorf(http.StatusText(http.StatusBadRequest)))
)

var (
	InvalidCursor = error.NewError(http.StatusNotFound, "ER9901", fmt.Errorf("cursor is invalid"))
)

var (
	InvalidRequest     = error.NewError(http.StatusBadRequest, "AU1001", fmt.Errorf("request is not valid"))
	UnparsableJSON     = error.NewError(http.StatusBadRequest, "AU1002", fmt.Errorf("failed on parsing JSON"))
	UnreadableResponse = error.NewError(http.StatusBadRequest, "AU1003", fmt.Errorf("failed to read response"))
)
