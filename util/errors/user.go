package errors

import (
	"fmt"
	"net/http"

	"gitlab.com/renodesper/gokit-microservices/util/error"
)

var (
	FailedPasswordMismatch = error.NewError(http.StatusBadRequest, "US1001", fmt.Errorf("password mismatch"))
)
