package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"gitlab.com/renodesper/gokit-microservices/service"
)

// Set ...
type Set struct {
	GetHealthCheckEndpoint endpoint.Endpoint
}

// New ...
func New(svc service.UserService, env string) Set {
	return Set{
		GetHealthCheckEndpoint: MakeHealthCheckEndpoint(svc),
	}
}
