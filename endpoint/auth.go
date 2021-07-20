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

	RegisterAuthRequest struct {
		Username    string `json:"username" validate:"required"`
		Email       string `json:"email" validate:"required"`
		Password    string `json:"password" validate:"required"`
		IsActive    bool   `json:"isActive"`
		IsDeleted   bool   `json:"isDeleted"`
		IsAdmin     bool   `json:"isAdmin"`
		CreatedFrom string `json:"createdFrom"`
	}

	RequestResetPasswordAuthRequest struct {
		Email string `validate:"required"`
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

func MakeRegisterAuthEndpoint(oauthSvc service.OauthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(RegisterAuthRequest)

		user, err := oauthSvc.Register(ctx, req.Username, req.Email, req.Password, req.IsActive, req.IsDeleted, req.IsAdmin, req.CreatedFrom)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}

func MakeRequestResetPasswordAuthEndpoint(oauthSvc service.OauthService, userSvc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(RequestResetPasswordAuthRequest)

		_, err = oauthSvc.RequestResetPassword(ctx, req.Email)
		if err != nil {
			return nil, err
		}

		return "OK", nil
	}
}
