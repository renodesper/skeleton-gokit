package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"gitlab.com/renodesper/gokit-microservices/service"
)

type Set struct {
	GoogleLoginAuthEndpoint          endpoint.Endpoint
	GoogleCallbackAuthEndpoint       endpoint.Endpoint
	LoginAuthEndpoint                endpoint.Endpoint
	LogoutAuthEndpoint               endpoint.Endpoint
	RegisterAuthEndpoint             endpoint.Endpoint
	VerifyRegistrationEndpoint       endpoint.Endpoint
	RequestResetPasswordAuthEndpoint endpoint.Endpoint
	VerifyResetPasswordEndpoint      endpoint.Endpoint
	GetHealthCheckEndpoint           endpoint.Endpoint
	CreateUserEndpoint               endpoint.Endpoint
	GetAllUsersEndpoint              endpoint.Endpoint
	GetUserEndpoint                  endpoint.Endpoint
	UpdateUserEndpoint               endpoint.Endpoint
	DeleteUserEndpoint               endpoint.Endpoint
	SetPasswordEndpoint              endpoint.Endpoint
	SetAccessTokenEndpoint           endpoint.Endpoint
	SetUserStatusEndpoint            endpoint.Endpoint
	SetUserRoleEndpoint              endpoint.Endpoint
	SetUserExpiryEndpoint            endpoint.Endpoint
}

// New ...
func New(
	env string,
	healthSvc service.HealthService,
	googleOauthSvc service.GoogleOauthService,
	oauthSvc service.OauthService,
	userSvc service.UserService,
	verificationSvc service.VerificationService,
) Set {
	return Set{
		GoogleLoginAuthEndpoint:          MakeGoogleLoginAuthEndpoint(googleOauthSvc),
		GoogleCallbackAuthEndpoint:       MakeGoogleCallbackAuthEndpoint(googleOauthSvc),
		LoginAuthEndpoint:                MakeLoginAuthEndpoint(oauthSvc),
		LogoutAuthEndpoint:               MakeLogoutAuthEndpoint(oauthSvc),
		RegisterAuthEndpoint:             MakeRegisterAuthEndpoint(oauthSvc),
		VerifyRegistrationEndpoint:       MakeVerifyRegistrationEndpoint(verificationSvc, userSvc),
		RequestResetPasswordAuthEndpoint: MakeRequestResetPasswordAuthEndpoint(oauthSvc, userSvc),
		VerifyResetPasswordEndpoint:      MakeVerifyResetPasswordEndpoint(verificationSvc, userSvc),
		GetHealthCheckEndpoint:           MakeHealthCheckEndpoint(healthSvc),
		CreateUserEndpoint:               MakeCreateUserEndpoint(userSvc),
		GetAllUsersEndpoint:              MakeGetAllUsersEndpoint(userSvc),
		GetUserEndpoint:                  MakeGetUserEndpoint(userSvc),
		UpdateUserEndpoint:               MakeUpdateUserEndpoint(userSvc),
		SetPasswordEndpoint:              MakeSetPasswordEndpoint(userSvc),
		SetAccessTokenEndpoint:           MakeSetAccessTokenEndpoint(userSvc),
		SetUserStatusEndpoint:            MakeSetUserStatusEndpoint(userSvc),
		SetUserRoleEndpoint:              MakeSetUserRoleEndpoint(userSvc),
		SetUserExpiryEndpoint:            MakeSetUserExpiryEndpoint(userSvc),
		DeleteUserEndpoint:               MakeDeleteUserEndpoint(userSvc),
	}
}
