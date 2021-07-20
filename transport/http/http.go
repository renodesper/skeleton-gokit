package http

import (
	"context"
	"net/http"

	kitendpoint "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-playground/validator"
	"github.com/go-zoo/bone"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
	"gitlab.com/renodesper/gokit-microservices/endpoint"
	m "gitlab.com/renodesper/gokit-microservices/middleware"
	"gitlab.com/renodesper/gokit-microservices/middleware/apiKey"
	"gitlab.com/renodesper/gokit-microservices/middleware/recover"
	ctxUtil "gitlab.com/renodesper/gokit-microservices/util/ctx"
	e "gitlab.com/renodesper/gokit-microservices/util/error"
	errs "gitlab.com/renodesper/gokit-microservices/util/errors"
	"gitlab.com/renodesper/gokit-microservices/util/logger"
	resp "gitlab.com/renodesper/gokit-microservices/util/response"
)

var (
	json     = jsoniter.ConfigCompatibleWithStandardLibrary
	validate *validator.Validate
)

// NewHTTPHandler ...
func NewHTTPHandler(endpoints endpoint.Set, log logger.Logger) http.Handler {
	r := bone.New()

	// NOTE: Will be executed on the HTTP request object before the request is decoded
	serverRequestOpts := []httptransport.RequestFunc{
		ctxUtil.ExtractRequestID,
		ctxUtil.ExtractApiKey,
		ctxUtil.ExtractJwtToken,
	}

	// NOTE: Will be executed on the HTTP response writer after the endpoint is invoked, but before anything written to the client
	serverResponseOpts := []httptransport.ServerResponseFunc{
		httptransport.SetContentType("application/vnd.api+json"),
	}

	serverOpts := []httptransport.ServerOption{
		httptransport.ServerBefore(serverRequestOpts...),
		httptransport.ServerAfter(serverResponseOpts...),
		httptransport.ServerErrorEncoder(encodeError),
	}

	// NOTE: Middlewares
	middlewares := m.Middlewares{
		Before: []kitendpoint.Middleware{
			recover.CreateMiddleware(log),
			apiKey.CreateMiddleware(log),
		},
		After: []kitendpoint.Middleware{},
	}
	publicMiddlewares := m.Middlewares{
		Before: []kitendpoint.Middleware{
			recover.CreateMiddleware(log),
		},
		After: []kitendpoint.Middleware{},
	}

	// NOTE: Routes
	r.NotFound(http.HandlerFunc(notFound))

	GoogleLoginAuthEndpoint := m.Chain(middlewares)(endpoints.GoogleLoginAuthEndpoint)
	r.Get("/auth/google", httptransport.NewServer(GoogleLoginAuthEndpoint, decodeNothing, encodeGoogleLoginAuthResponse, serverOpts...))

	GoogleCallbackAuthEndpoint := m.Chain(middlewares)(endpoints.GoogleCallbackAuthEndpoint)
	r.Get("/auth/google/callback", httptransport.NewServer(GoogleCallbackAuthEndpoint, decodeGoogleCallbackAuthRequest, encodeResponse, serverOpts...))

	LoginAuthEndpoint := m.Chain(middlewares)(endpoints.LoginAuthEndpoint)
	r.Post("/login", httptransport.NewServer(LoginAuthEndpoint, decodeLoginAuthRequest, encodeResponse, serverOpts...))

	LogoutAuthEndpoint := m.Chain(middlewares)(endpoints.LogoutAuthEndpoint)
	r.Post("/logout", httptransport.NewServer(LogoutAuthEndpoint, decodeLogoutAuthRequest, encodeResponse, serverOpts...))

	RegisterAuthEndpoint := m.Chain(middlewares)(endpoints.RegisterAuthEndpoint)
	r.Post("/register", httptransport.NewServer(RegisterAuthEndpoint, decodeRegisterAuthRequest, encodeResponse, serverOpts...))

	VerifyRegistrationEndpoint := m.Chain(publicMiddlewares)(endpoints.VerifyRegistrationEndpoint)
	r.Get("/confirm/:token", httptransport.NewServer(VerifyRegistrationEndpoint, decodeVerifyTokenRequest, encodeVerifyRegistrationResponse, serverOpts...))

	RequestResetPasswordAuthEndpoint := m.Chain(middlewares)(endpoints.RequestResetPasswordAuthEndpoint)
	r.Post("/reset-password", httptransport.NewServer(RequestResetPasswordAuthEndpoint, decodeRequestResetPasswordAuthRequest, encodeResponse, serverOpts...))

	VerifyResetPasswordEndpoint := m.Chain(publicMiddlewares)(endpoints.VerifyResetPasswordEndpoint)
	r.Get("/reset-password/:token", httptransport.NewServer(VerifyResetPasswordEndpoint, decodeVerifyTokenRequest, encodeVerifyResetPasswordResponse, serverOpts...))

	GetHealthCheckEndpoint := m.Chain(middlewares)(endpoints.GetHealthCheckEndpoint)
	r.Get("/health", httptransport.NewServer(GetHealthCheckEndpoint, decodeNothing, encodeResponse, serverOpts...))

	CreateUserEndpoint := m.Chain(middlewares)(endpoints.CreateUserEndpoint)
	r.Post("/users", httptransport.NewServer(CreateUserEndpoint, decodeCreateUserRequest, encodeResponse, serverOpts...))

	GetAllUsersEndpoint := m.Chain(middlewares)(endpoints.GetAllUsersEndpoint)
	r.Get("/users", httptransport.NewServer(GetAllUsersEndpoint, decodeGetAllUsersRequest, encodeResponse, serverOpts...))

	GetUserEndpoint := m.Chain(middlewares)(endpoints.GetUserEndpoint)
	r.Get("/users/:id", httptransport.NewServer(GetUserEndpoint, decodeGetUserRequest, encodeResponse, serverOpts...))

	UpdateUserEndpoint := m.Chain(middlewares)(endpoints.UpdateUserEndpoint)
	r.Put("/users/:id", httptransport.NewServer(UpdateUserEndpoint, decodeUpdateUserRequest, encodeResponse, serverOpts...))

	SetAccessTokenEndpoint := m.Chain(middlewares)(endpoints.SetAccessTokenEndpoint)
	r.Put("/users/:id/accessToken", httptransport.NewServer(SetAccessTokenEndpoint, decodeSetAccessTokenRequest, encodeResponse, serverOpts...))

	SetUserStatusEndpoint := m.Chain(middlewares)(endpoints.SetUserStatusEndpoint)
	r.Put("/users/:id/status", httptransport.NewServer(SetUserStatusEndpoint, decodeSetUserStatusRequest, encodeResponse, serverOpts...))

	SetUserRoleEndpoint := m.Chain(middlewares)(endpoints.SetUserRoleEndpoint)
	r.Put("/users/:id/role", httptransport.NewServer(SetUserRoleEndpoint, decodeSetUserRoleRequest, encodeResponse, serverOpts...))

	SetUserExpiryEndpoint := m.Chain(middlewares)(endpoints.SetUserExpiryEndpoint)
	r.Put("/users/:id/expiry", httptransport.NewServer(SetUserExpiryEndpoint, decodeUpdateUserRequest, encodeResponse, serverOpts...))

	DeleteUserEndpoint := m.Chain(middlewares)(endpoints.DeleteUserEndpoint)
	r.Delete("/users/:id", httptransport.NewServer(DeleteUserEndpoint, decodeDeleteUserRequest, encodeResponse, serverOpts...))

	return r
}

// decodeNothing returns (nil, nil) as placeholder for httptransport.DecodeRequestFunc
func decodeNothing(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	requestID := ctxUtil.GetRequestID(ctx)
	return json.NewEncoder(w).Encode(&resp.SuccessResponse{
		Data: response,
		Meta: resp.PopulateMeta(requestID),
	})
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	er := err.(e.Error)
	status := er.Status
	if status == 0 {
		status = 500
	}

	if viper.GetString("app.env") == "production" {
		er = er.WithoutStackTrace()
	}

	requestID := ctxUtil.GetRequestID(ctx)

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&resp.ErrorResponse{
		Errors: []e.Error{er},
		Meta:   resp.PopulateMeta(requestID),
	})
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&resp.ErrorResponse{
		Errors: []e.Error{errs.StatusNotFound.WithoutStackTrace()},
		Meta:   resp.PopulateMeta(r.Header.Get("X-Request-Id")),
	})
}
