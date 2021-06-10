package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"gitlab.com/renodesper/gokit-microservices/service"
	authUtil "gitlab.com/renodesper/gokit-microservices/util/auth"
)

type (
	CallbackAuthRequest struct {
		Code string `json:"code" validate:"required"`
	}

	LoginAuthRequest struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	LogoutAuthRequest struct {
		UserID string `json:"userId" validate:"required"`
	}
)

func MakeGoogleLoginAuthEndpoint(googleOauthSvc service.GoogleOauthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		config := authUtil.GetGoogleOauthConfig()
		return config, nil
	}
}

func MakeGoogleCallbackAuthEndpoint(googleOauthSvc service.GoogleOauthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CallbackAuthRequest)

		token, err := googleOauthSvc.OauthCallback(ctx, req.Code)
		if err != nil {
			return nil, err
		}

		return token, nil
	}
}

func MakeLoginAuthEndpoint(oauthSvc service.OauthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(LoginAuthRequest)

		token, err := oauthSvc.Login(ctx, req.Email, req.Password)
		if err != nil {
			return nil, err
		}

		return token, nil
	}
}

func MakeLogoutAuthEndpoint(oauthSvc service.OauthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(LogoutAuthRequest)

		userID, err := uuid.Parse(req.UserID)
		if err != nil {
			return nil, err
		}

		err = oauthSvc.Logout(ctx, userID)
		if err != nil {
			return nil, err
		}

		return "OK", nil
	}
}
