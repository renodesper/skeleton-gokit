package jwt

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	authUtil "gitlab.com/renodesper/gokit-microservices/util/auth"
	ctxUtil "gitlab.com/renodesper/gokit-microservices/util/ctx"
	"gitlab.com/renodesper/gokit-microservices/util/logger"
)

// CreateMiddleware ...
func CreateMiddleware(log logger.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			jwtToken := ctxUtil.GetJwtToken(ctx)

			// NOTE: identity is the id of user
			identity, err := authUtil.ParseJWTWithClaims(jwtToken)
			if err != nil {
				return nil, err
			}

			log.Info(identity)

			// TODO: Check user existence in DB
			// ...

			// TODO: Store identity into context
			// ...

			response, err := next(ctx, request)
			return response, err
		}
	}
}
