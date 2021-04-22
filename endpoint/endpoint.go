package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"gitlab.com/renodesper/gokit-microservices/service"
)

type Set struct {
	GetHealthCheckEndpoint endpoint.Endpoint
	CreateUserEndpoint     endpoint.Endpoint
	GetAllUsersEndpoint    endpoint.Endpoint
	GetUserEndpoint        endpoint.Endpoint
}

// New ...
func New(healthSvc service.HealthService, userSvc service.UserService, env string) Set {
	return Set{
		GetHealthCheckEndpoint: MakeHealthCheckEndpoint(healthSvc),
		CreateUserEndpoint:     MakeCreateUserEndpoint(userSvc),
		GetAllUsersEndpoint:    MakeGetAllUsersEndpoint(userSvc),
		GetUserEndpoint:        MakeGetUserEndpoint(userSvc),
	}
}
