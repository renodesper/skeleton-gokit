package request

import (
	"context"
	"encoding/json"

	"github.com/go-kit/kit/endpoint"
	"gitlab.com/renodesper/gokit-microservices/util/logger"
)

// Request ...
func Request(log logger.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			b, err := json.Marshal(request)
			if err != nil {
				return nil, err
			}

			log.Info(string(b))

			response, err := next(ctx, request)
			return response, err
		}
	}
}
