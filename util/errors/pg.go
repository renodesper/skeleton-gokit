package errors

import (
	"fmt"
	"net/http"

	"gitlab.com/renodesper/gokit-microservices/util/error"
)

var (
	FailedNoRows = error.NewError(http.StatusBadRequest, "PG3000", fmt.Errorf("no rows selected"))
)

var (
	FailedUserExist    = error.NewError(http.StatusBadRequest, "PG3001", fmt.Errorf("user is already exist"))
	FailedUserNotFound = error.NewError(http.StatusNotFound, "PG3002", fmt.Errorf("user cannot be found"))
	FailedUserCreate   = error.NewError(http.StatusInternalServerError, "PG3003", fmt.Errorf("unable to create user"))
	FailedUserUpdate   = error.NewError(http.StatusInternalServerError, "PG3004", fmt.Errorf("unable to update user"))
	FailedUserDelete   = error.NewError(http.StatusInternalServerError, "PG3005", fmt.Errorf("unable to delete user"))
	FailedUserFetch    = error.NewError(http.StatusInternalServerError, "PG3006", fmt.Errorf("unable to fetch user"))
	FailedUsersFetch   = error.NewError(http.StatusInternalServerError, "PG3007", fmt.Errorf("unable to fetch users"))
)

var (
	FailedEmailExist    = error.NewError(http.StatusBadRequest, "PG3101", fmt.Errorf("email is already exist"))
	FailedUsernameExist = error.NewError(http.StatusBadRequest, "PG3102", fmt.Errorf("username is already exist"))
)
