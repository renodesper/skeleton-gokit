package errors

import (
	"fmt"
	"net/http"

	"gitlab.com/renodesper/gokit-microservices/util/error"
)

var (
	FailedNoRows        = error.NewError(http.StatusBadRequest, "PG3000", fmt.Errorf("no rows selected"))
	FailedEmailExist    = error.NewError(http.StatusBadRequest, "PG3001", fmt.Errorf("email is already exist"))
	FailedUsernameExist = error.NewError(http.StatusBadRequest, "PG3002", fmt.Errorf("username is already exist"))
)

var (
	FailedUserExist    = error.NewError(http.StatusBadRequest, "PG3101", fmt.Errorf("user is already exist"))
	FailedUserNotFound = error.NewError(http.StatusNotFound, "PG3102", fmt.Errorf("user cannot be found"))
	FailedUserCreate   = error.NewError(http.StatusInternalServerError, "PG3103", fmt.Errorf("unable to create user"))
	FailedUserUpdate   = error.NewError(http.StatusInternalServerError, "PG3104", fmt.Errorf("unable to update user"))
	FailedUserDelete   = error.NewError(http.StatusInternalServerError, "PG3105", fmt.Errorf("unable to delete user"))
	FailedUserFetch    = error.NewError(http.StatusInternalServerError, "PG3106", fmt.Errorf("unable to fetch user"))
	FailedUsersFetch   = error.NewError(http.StatusInternalServerError, "PG3107", fmt.Errorf("unable to fetch users"))
)

var (
	FailedVerificationCreate = error.NewError(http.StatusInternalServerError, "PG3201", fmt.Errorf("unable to create verification"))
)
