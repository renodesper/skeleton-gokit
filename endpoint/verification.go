package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"gitlab.com/renodesper/gokit-microservices/service"
)

type (
	VerifyRegistrationRequest struct {
		Token string
	}
)

// MakeVerifyRegistrationEndpoint ...
func MakeVerifyRegistrationEndpoint(svc service.VerificationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(VerifyRegistrationRequest)

		err = svc.VerifyRegistration(ctx, req.Token)
		if err != nil {
			return nil, err
		}

		return "OK", nil
	}
}
