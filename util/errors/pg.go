package errors

import (
	"fmt"
	"net/http"

	"gitlab.com/renodesper/gokit-microservices/util/error"
)

var (
	FailedUserExist    = error.NewError(http.StatusBadRequest, "PG3001", fmt.Errorf("User is already exist"))
	FailedUserNotFound = error.NewError(http.StatusNotFound, "PG3002", fmt.Errorf("User cannot be found"))
	FailedUserCreate   = error.NewError(http.StatusInternalServerError, "PG3003", fmt.Errorf("Unable to create user"))
	FailedUserUpdate   = error.NewError(http.StatusInternalServerError, "PG3004", fmt.Errorf("Unable to update user"))
	FailedUserDelete   = error.NewError(http.StatusInternalServerError, "PG3005", fmt.Errorf("Unable to delete user"))
	FailedUserFetch    = error.NewError(http.StatusInternalServerError, "PG3006", fmt.Errorf("Unable to fetch user"))
	FailedUsersFetch   = error.NewError(http.StatusInternalServerError, "PG3007", fmt.Errorf("Unable to fetch users"))
)

var (
	FailedEmailExist    = error.NewError(http.StatusBadRequest, "PG3101", fmt.Errorf("Email is already exist"))
	FailedUsernameExist = error.NewError(http.StatusBadRequest, "PG3102", fmt.Errorf("Username is already exist"))
)
