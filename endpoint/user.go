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
		ID uuid.UUID
	}

	UpdateUserRequest struct {
		ID       uuid.UUID
		Username string
		Email    string
		Password string
	}

	DeleteUserRequest struct {
		ID uuid.UUID
	}

	SetAccessTokenRequest struct {
		ID           uuid.UUID
		AccessToken  string
		RefreshToken string
		ExpiredAt    time.Time
	}

	SetUserStatusRequest struct {
		ID       uuid.UUID
		IsActive bool
	}

	SetUserRoleRequest struct {
		ID      uuid.UUID
		IsAdmin bool
	}

	SetUserExpiryRequest struct {
		ID        uuid.UUID
		ExpiredAt time.Time
	}
)

func MakeCreateUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateUserRequest)

		userReq := service.CreateUserRequest{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
			IsAdmin:  req.IsAdmin,
		}

		isSuccess, err := svc.CreateUser(ctx, &userReq)
		if err != nil {
			return nil, err
		}

		return isSuccess, nil
	}
}

func MakeGetAllUsersEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetAllUsersRequest)

		users, err := svc.GetAllUsers(ctx, req.SortBy, req.Sort, req.Skip, req.Limit)
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

func MakeGetUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetUserRequest)

		users, err := svc.GetUser(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		return users, nil
	}
}

func MakeUpdateUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateUserRequest)

		userPayload := service.UpdateUserRequest{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
		}

		user, err := svc.UpdateUser(ctx, req.ID, &userPayload)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}

func MakeSetAccessTokenEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SetAccessTokenRequest)

		user, err := svc.SetAccessToken(ctx, req.ID, req.AccessToken, req.RefreshToken, req.ExpiredAt)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}

func MakeSetUserStatusEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SetUserStatusRequest)

		user, err := svc.SetUserStatus(ctx, req.ID, req.IsActive)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}

func MakeSetUserRoleEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SetUserRoleRequest)

		user, err := svc.SetUserRole(ctx, req.ID, req.IsAdmin)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}

func MakeSetUserExpiryEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SetUserExpiryRequest)

		user, err := svc.SetUserExpiry(ctx, req.ID, req.ExpiredAt)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}

func MakeDeleteUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteUserRequest)

		user, err := svc.DeleteUser(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}
