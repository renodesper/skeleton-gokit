package errors

import (
	"fmt"
	"net/http"

	"gitlab.com/renodesper/gokit-microservices/util/error"
)

var (
	// AU1000 returned when client try to access undefined endpoint
	AU1000 = error.NewError(http.StatusNotFound, "AU1000", fmt.Errorf(http.StatusText(http.StatusNotFound)))

	// AU1001 returned when ...
	AU1001 = error.NewError(http.StatusBadRequest, "AU1001", fmt.Errorf(http.StatusText(http.StatusBadRequest)))

	// AU1002 returned when ...
	AU1002 = error.NewError(http.StatusBadRequest, "AU1002", fmt.Errorf("Request is not valid"))

	// AU1003 returned when ...
	AU1003 = error.NewError(http.StatusBadRequest, "AU1003", fmt.Errorf("Failed on parsing JSON"))
)

var (
	// AU2000 returned when session token is invalid
	AU2000 = error.NewError(401, "AU2000", fmt.Errorf("Session token is invalid"))
)
