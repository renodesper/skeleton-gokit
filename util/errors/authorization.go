package errors

import (
	"fmt"
	"net/http"

	"gitlab.com/renodesper/gokit-microservices/util/error"
)

var (
	InvalidSessionToken    = error.NewError(http.StatusUnauthorized, "AU2001", fmt.Errorf("session token is invalid"))
	InvalidJwtToken        = error.NewError(http.StatusUnauthorized, "AU2002", fmt.Errorf("jwt token is invalid"))
	InvalidApiKey          = error.NewError(http.StatusUnauthorized, "AU2003", fmt.Errorf("api key is invalid"))
	InvalidLoginCredential = error.NewError(http.StatusBadRequest, "AU3004", fmt.Errorf("invalid login credential"))
)

var (
	InvalidGoogleOauthState        = error.NewError(http.StatusBadRequest, "AU3001", fmt.Errorf("google oauthState is invalid"))
	InvalidGoogleOauthCodeExchange = error.NewError(http.StatusBadRequest, "AU3002", fmt.Errorf("google oauth code exchange is invalid"))
	FailedGoogleUserFetch          = error.NewError(http.StatusInternalServerError, "AU3003", fmt.Errorf("unable to fetch user"))
)
