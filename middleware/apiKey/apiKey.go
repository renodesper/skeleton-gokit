package apiKey

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/spf13/viper"
	ctxUtil "gitlab.com/renodesper/gokit-microservices/util/ctx"
	"gitlab.com/renodesper/gokit-microservices/util/errors"
	"gitlab.com/renodesper/gokit-microservices/util/logger"
)

// CreateMiddleware ...
func CreateMiddleware(log logger.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			apiKey := ctxUtil.GetApiKey(ctx)

			// NOTE: We can store API Key in another service like consul, so each client will have its own API Key
			API_KEY := viper.GetString("app.api_key")
			if apiKey != API_KEY {
				return nil, errors.InvalidApiKey
			}

			response, err := next(ctx, request)
			return response, err
		}
	}
}
