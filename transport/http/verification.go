package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-zoo/bone"
	"gitlab.com/renodesper/gokit-microservices/endpoint"
)

func decodeVerifyTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.VerifyTokenRequest

	token := bone.GetValue(r, "token")
	req.Token = token

	return req, nil
}

func encodeVerifyRegistrationResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
	return nil
}

func encodeVerifyResetPasswordResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	tokenResponse := response.(endpoint.VerifyTokenResponse)

	// NOTE: Redirect to update password page
	uri := fmt.Sprintf("/update-password/%s", tokenResponse.Token)

	w.Header().Set("Location", uri)
	w.WriteHeader(http.StatusSeeOther)
	return nil
}
