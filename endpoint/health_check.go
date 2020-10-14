package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"gitlab.com/renodesper/gokit-microservices/service"
)

// HealthCheckResponse ...
type HealthCheckResponse struct {
	Version string `json:"version"`
}

// MakeHealthCheckEndpoint ...
func MakeHealthCheckEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return HealthCheckResponse{
			Version: svc.HealthCheck(),
		}, nil
	}
}
