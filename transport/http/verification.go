package http

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/go-zoo/bone"
	"gitlab.com/renodesper/gokit-microservices/endpoint"
	"gitlab.com/renodesper/gokit-microservices/util/errors"
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

func decodeVerifyResetPasswordAuthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.VerifyResetPasswordRequest

	token := bone.GetValue(r, "token")
	req.Token = token

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.UnparsableJSON
	}
	defer r.Body.Close()

	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, errors.InvalidRequest
	}

	return req, nil
}

func encodeVerifyResetPasswordResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
	return nil
}
