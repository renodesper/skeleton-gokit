package http

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-playground/validator"
	"github.com/go-zoo/bone"
	"gitlab.com/renodesper/gokit-microservices/pkg/endpoint"
	e "gitlab.com/renodesper/gokit-microservices/pkg/util/error"
)

type (
	// ErrorResponse ...
	ErrorResponse struct {
		Error []e.Error `json:"errors"`
	}

	// SuccessResponse ...
	SuccessResponse struct {
		Data interface{} `json:"data"`
	}

	meta map[string]interface{}
)

var (
	validate *validator.Validate
)

// MakeHTTPHandler ...
func MakeHTTPHandler(endpoints endpoint.Set) http.Handler {
	r := bone.New()

	serverRequestOpts := []httptransport.RequestFunc{
		// NOTE: We can put a function that receive context and return context here
	}

	serverResponseOpts := []httptransport.ServerResponseFunc{
		httptransport.SetContentType("application/vnd.api+json"),
	}

	serverOpts := []httptransport.ServerOption{
		httptransport.ServerBefore(serverRequestOpts...),
		httptransport.ServerAfter(serverResponseOpts...),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.NotFound(http.HandlerFunc(notFound))

	r.Get("/health", httptransport.NewServer(
		endpoints.GetHealthCheckEndpoint,
		decodeHealthCheckRequest,
		encodeResponse,
		serverOpts...,
	))

	return r
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(&SuccessResponse{
		Data: response,
	})
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	// TODO: Parse from "err" param
	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>
	statusCode := http.StatusBadRequest
	errors := []e.Error{
		e.Error{
			Status: http.StatusBadRequest,
			Code:   "XXX-123",
		},
	}
	// <<<<<<<<<<<<<<<<<<<<<<<<<<<<

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(&ErrorResponse{Error: errors})
}

func decodeAndValidate(r *http.Request, model interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		return e.Error{
			Status:  http.StatusBadRequest,
			Code:    "XXX-123",
			Message: "Failed on parsing JSON",
		}
	}
	defer r.Body.Close()

	validate = validator.New()
	if err := validate.Struct(model); err != nil {
		return e.Error{
			Status:  http.StatusBadRequest,
			Code:    "XXX-123",
			Message: "Request is not valid",
		}
	}

	return nil
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(&ErrorResponse{
		Error: []e.Error{
			e.Error{
				Status: http.StatusNotFound,
				Code:   "XXX-123",
			},
		},
	})
}
