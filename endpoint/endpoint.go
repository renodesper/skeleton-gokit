package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"gitlab.com/renodesper/gokit-microservices/service"
)

type Set struct {
	GetLoginAuthEndpoint    endpoint.Endpoint
	GetCallbackAuthEndpoint endpoint.Endpoint
	GetLogoutAuthEndpoint   endpoint.Endpoint
	GetHealthCheckEndpoint  endpoint.Endpoint
	CreateUserEndpoint      endpoint.Endpoint
	GetAllUsersEndpoint     endpoint.Endpoint
	GetUserEndpoint         endpoint.Endpoint
	UpdateUserEndpoint      endpoint.Endpoint
	DeleteUserEndpoint      endpoint.Endpoint
	SetAccessTokenEndpoint  endpoint.Endpoint
	SetUserStatusEndpoint   endpoint.Endpoint
	SetUserRoleEndpoint     endpoint.Endpoint
	SetUserExpiryEndpoint   endpoint.Endpoint
}

// New ...
func New(healthSvc service.HealthService, googleOauthSvc service.GoogleOauthService, userSvc service.UserService, env string) Set {
	return Set{
		GetLoginAuthEndpoint:    MakeLoginAuthEndpoint(googleOauthSvc),
		GetCallbackAuthEndpoint: MakeCallbackAuthEndpoint(googleOauthSvc),
		GetLogoutAuthEndpoint:   MakeLogoutAuthEndpoint(googleOauthSvc),
		GetHealthCheckEndpoint:  MakeHealthCheckEndpoint(healthSvc),
		CreateUserEndpoint:      MakeCreateUserEndpoint(userSvc),
		GetAllUsersEndpoint:     MakeGetAllUsersEndpoint(userSvc),
		GetUserEndpoint:         MakeGetUserEndpoint(userSvc),
		UpdateUserEndpoint:      MakeUpdateUserEndpoint(userSvc),
		SetAccessTokenEndpoint:  MakeSetAccessTokenEndpoint(userSvc),
		SetUserStatusEndpoint:   MakeSetUserStatusEndpoint(userSvc),
		SetUserRoleEndpoint:     MakeSetUserRoleEndpoint(userSvc),
		SetUserExpiryEndpoint:   MakeSetUserExpiryEndpoint(userSvc),
		DeleteUserEndpoint:      MakeDeleteUserEndpoint(userSvc),
	}
}
