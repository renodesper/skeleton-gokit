package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"gitlab.com/renodesper/gokit-microservices/service"
	authUtil "gitlab.com/renodesper/gokit-microservices/util/auth"
)

type (
	CallbackAuthRequest struct {
		Code string `json:"code" validate:"required"`
	}
)

func MakeLoginAuthEndpoint(googleOauthSvc service.GoogleOauthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		config := authUtil.GetGoogleOauthConfig()
		return config, nil
	}
}

func MakeCallbackAuthEndpoint(googleOauthSvc service.GoogleOauthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CallbackAuthRequest)

		userData, err := googleOauthSvc.OauthCallback(ctx, req.Code)
		if err != nil {
			return nil, err
		}

		return userData, nil
	}
}

// TODO:
func MakeLogoutAuthEndpoint(googleOauthSvc service.GoogleOauthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return nil, nil
	}
}
