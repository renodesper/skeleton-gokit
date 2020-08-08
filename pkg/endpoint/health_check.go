package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"gitlab.com/renodesper/gokit-microservices/pkg/service"
)

// MakeHealthCheckEndpoint ...
func MakeHealthCheckEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.HealthCheck(), nil
	}
}
