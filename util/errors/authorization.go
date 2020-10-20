package errors

import (
	"fmt"
	"net/http"

	"gitlab.com/renodesper/gokit-microservices/util/error"
)

var (
	// InvalidRequest returned when ...
	InvalidRequest = error.NewError(http.StatusBadRequest, "AU1001", fmt.Errorf("Request is not valid"))

	// UnparsableJSON returned when ...
	UnparsableJSON = error.NewError(http.StatusBadRequest, "AU1002", fmt.Errorf("Failed on parsing JSON"))
)

var (
	// InvalidSessionToken returned when session token is invalid
	InvalidSessionToken = error.NewError(http.StatusUnauthorized, "AU2001", fmt.Errorf("Session token is invalid"))
)
