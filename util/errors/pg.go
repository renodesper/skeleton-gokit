package errors

import (
	"fmt"
	"net/http"

	"gitlab.com/renodesper/gokit-microservices/util/error"
)

var (
	// FailedUserCreate ...
	FailedUserCreate = error.NewError(http.StatusInternalServerError, "PG3001", fmt.Errorf("Unable to create user"))
)
