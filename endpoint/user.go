package endpoint

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"gitlab.com/renodesper/gokit-microservices/service"
)

type (
	CreateUserRequest struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
		IsAdmin  bool   `json:"isAdmin" validate:"required"`
	}

	GetAllUsersRequest struct {
		SortBy string `json:"sortBy"`
		Sort   string `json:"sort"`
		Skip   int    `json:"skip"`
		Limit  int    `json:"limit"`
	}

	GetUserRequest struct {
		ID uuid.UUID `json:"id" validate:"required"`
	}

	UpdateUserRequest struct {
		ID       uuid.UUID `json:"id" validate:"required"`
		Username string
		Email    string
		Password string
	}

	DeleteUserRequest struct {
		ID uuid.UUID `json:"id" validate:"required"`
	}

	SetPasswordRequest struct {
		ID             uuid.UUID `json:"id" validate:"required"`
		Password       string    `json:"password" validate:"required"`
		VerifyPassword string    `json:"verifyPassword" validate:"required"`
	}

	SetAccessTokenRequest struct {
		ID           uuid.UUID `json:"id" validate:"required"`
		AccessToken  string    `json:"accessToken" validate:"required"`
		RefreshToken string    `json:"refreshToken" validate:"required"`
		ExpiredAt    time.Time
	}

	SetUserStatusRequest struct {
		ID       uuid.UUID `json:"id" validate:"required"`
		IsActive bool      `json:"isActive" validate:"required"`
	}

	SetUserRoleRequest struct {
		ID      uuid.UUID `json:"id" validate:"required"`
		IsAdmin bool      `json:"isAdmin" validate:"required"`
	}

	SetUserExpiryRequest struct {
		ID        uuid.UUID `json:"id" validate:"required"`
		ExpiredAt time.Time `json:"expiredAt" validate:"required"`
	}
)

func MakeCreateUserEndpoint(userSvc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateUserRequest)

		userReq := service.CreateUserRequest{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
			IsAdmin:  req.IsAdmin,
		}

		isSuccess, err := userSvc.CreateUser(ctx, &userReq)
		if err != nil {
			return nil, err
		}

		return isSuccess, nil
	}
}

func MakeGetAllUsersEndpoint(userSvc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetAllUsersRequest)

		users, err := userSvc.GetAllUsers(ctx, req.SortBy, req.Sort, req.Skip, req.Limit)
		if err != nil {
			return nil, err
		}

		response = map[string]interface{}{
			"users": users,
			"pagination": map[string]interface{}{
				"sortBy": req.SortBy,
				"sort":   req.Sort,
				"skip":   req.Skip,
				"limit":  req.Limit,
			},
		}

		return response, nil
	}
}

func MakeGetUserEndpoint(userSvc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetUserRequest)

		user, err := userSvc.GetUser(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}

func MakeUpdateUserEndpoint(userSvc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateUserRequest)

		userPayload := service.UpdateUserRequest{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
		}

		user, err := userSvc.UpdateUser(ctx, req.ID, &userPayload)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}

func MakeSetPasswordEndpoint(userSvc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SetPasswordRequest)

		user, err := userSvc.SetPassword(ctx, req.ID, req.Password, req.VerifyPassword)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}

func MakeSetAccessTokenEndpoint(userSvc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SetAccessTokenRequest)

		user, err := userSvc.SetAccessToken(ctx, req.ID, req.AccessToken, req.RefreshToken, req.ExpiredAt)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}

func MakeSetUserStatusEndpoint(userSvc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SetUserStatusRequest)

		user, err := userSvc.SetUserStatus(ctx, req.ID, req.IsActive)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}

func MakeSetUserRoleEndpoint(userSvc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SetUserRoleRequest)

		user, err := userSvc.SetUserRole(ctx, req.ID, req.IsAdmin)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}

func MakeSetUserExpiryEndpoint(userSvc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SetUserExpiryRequest)

		user, err := userSvc.SetUserExpiry(ctx, req.ID, req.ExpiredAt)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}

func MakeDeleteUserEndpoint(userSvc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteUserRequest)

		user, err := userSvc.DeleteUser(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}
