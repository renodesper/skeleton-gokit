package errors

import (
	"fmt"
	"net/http"

	"gitlab.com/renodesper/gokit-microservices/util/error"
)

var (
	// InvalidSessionToken returned when session token is invalid
	InvalidSessionToken = error.NewError(http.StatusUnauthorized, "AU2001", fmt.Errorf("Session token is invalid"))

	// InvalidLoginCredential returned when login credential (email or password) is invalid
	InvalidLoginCredential = error.NewError(http.StatusBadRequest, "AU3004", fmt.Errorf("Invalid login credential"))
)

var (
	// InvalidGoogleOauthState returned when Google oauthState is invalid
	InvalidGoogleOauthState = error.NewError(http.StatusBadRequest, "AU3001", fmt.Errorf("Google oauthState is invalid"))

	// InvalidGoogleOauthCodeExchange returned when Google code exchange is invalid
	InvalidGoogleOauthCodeExchange = error.NewError(http.StatusBadRequest, "AU3002", fmt.Errorf("Google oauth code exchange is invalid"))

	// FailedGoogleUserFetch returned when failed in fetching user data
	FailedGoogleUserFetch = error.NewError(http.StatusInternalServerError, "AU3003", fmt.Errorf("Unable to fetch user"))
)
