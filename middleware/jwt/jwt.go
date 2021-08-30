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

			// NOTE: the first return valus is identity which is the id of user
			_, err := authUtil.ParseJWTWithClaims(jwtToken)
			if err != nil {
				return nil, err
			}

			// NOTE: Store identity into context if needed
			response, err := next(ctx, request)
			return response, err
		}
	}
}
