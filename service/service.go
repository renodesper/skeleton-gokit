package service

type (
	// UserService ...
	UserService interface {
		HealthCheck() string
	}

	// UserSvc ...
	UserSvc struct {
		// NOTE: repository that the service needs
		// ExampleAPI    *example.ExampleAPI
	}
)

// New creates user service
func New() UserService {
	return &UserSvc{
		// ExampleAPI: example,
	}
}
