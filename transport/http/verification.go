package http

import (
	"context"
	"net/http"

	"github.com/go-zoo/bone"
	"gitlab.com/renodesper/gokit-microservices/endpoint"
)

func decodeVerifyRegistrationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.VerifyRegistrationRequest

	token := bone.GetValue(r, "token")
	req.Token = token

	return req, nil
}
