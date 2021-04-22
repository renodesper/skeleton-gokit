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
	"gitlab.com/renodesper/gokit-microservices/middleware/recover"
	ctxutil "gitlab.com/renodesper/gokit-microservices/util/ctx"
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
		ctxutil.FromHTTPRequest,
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

	r.NotFound(http.HandlerFunc(notFound))

	// NOTE: Empty middlewares for the sake of example
	middlewares := m.Middlewares{
		Before: []kitendpoint.Middleware{
			recover.CreateMiddleware(),
		},
		After: []kitendpoint.Middleware{},
	}

	GetHealthCheckEndpoint := m.Chain(middlewares)(endpoints.GetHealthCheckEndpoint)
	r.Get("/health", httptransport.NewServer(
		GetHealthCheckEndpoint,
		decodeNothing,
		encodeResponse,
		serverOpts...,
	))

	CreateUserEndpoint := m.Chain(middlewares)(endpoints.CreateUserEndpoint)
	r.Post("/users", httptransport.NewServer(
		CreateUserEndpoint,
		decodeCreateUserRequest,
		encodeResponse,
		serverOpts...,
	))

	GetAllUsersEndpoint := m.Chain(middlewares)(endpoints.GetAllUsersEndpoint)
	r.Get("/users", httptransport.NewServer(
		GetAllUsersEndpoint,
		decodeGetAllUsersRequest,
		encodeResponse,
		serverOpts...,
	))

	GetUserEndpoint := m.Chain(middlewares)(endpoints.GetUserEndpoint)
	r.Get("/users/:id", httptransport.NewServer(
		GetUserEndpoint,
		decodeGetUserRequest,
		encodeResponse,
		serverOpts...,
	))

	return r
}

// decodeNothing returns (nil, nil) as placeholder for httptransport.DecodeRequestFunc
func decodeNothing(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	requestID := ctxutil.GetRequestID(ctx)
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

	requestID := ctxutil.GetRequestID(ctx)

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
