package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"gitlab.com/renodesper/gokit-microservices/service"
)

type (
	VerifyTokenRequest struct {
		Token string
	}

	VerifyTokenResponse struct {
		Token string
	}
)

// MakeVerifyRegistrationEndpoint ...
func MakeVerifyRegistrationEndpoint(verificationSvc service.VerificationService, userSvc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(VerifyTokenRequest)

		userID, err := verificationSvc.VerifyRegistration(ctx, req.Token)
		if err != nil {
			return nil, err
		}

		isActive := true
		_, err = userSvc.SetUserStatus(ctx, userID, isActive)
		if err != nil {
			return nil, err
		}

		return "OK", nil
	}
}

// MakeVerifyResetPasswordEndpoint ...
func MakeVerifyResetPasswordEndpoint(verificationSvc service.VerificationService, userSvc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(VerifyTokenRequest)

		token, err := verificationSvc.VerifyResetPassword(ctx, req.Token)
		if err != nil {
			return nil, err
		}

		tokenResponse := VerifyTokenResponse{
			Token: token.String(),
		}

		return tokenResponse, nil
	}
}
