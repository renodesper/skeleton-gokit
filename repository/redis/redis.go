package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"gitlab.com/renodesper/gokit-microservices/repository"
	utilRedis "gitlab.com/renodesper/gokit-microservices/util/connection/redis"
)

type (
	// RedisService ...
	RedisService interface {
		SetUser(ctx context.Context, user repository.User) error
		GetUserByID(ctx context.Context, userID string) (repository.User, error)
	}

	redisService struct {
		Client *redis.Client
	}
)

// SetUser ...
func (r *redisService) SetUser(ctx context.Context, user repository.User) error {
	redisKey := fmt.Sprintf("user:%s", user.ID.String())
	err := r.Client.Set(ctx, redisKey, &user, 0)
	if err != nil {
		return err.Err()
	}
	return nil
}

// GetUserByID ...
func (r *redisService) GetUserByID(ctx context.Context, userID string) (repository.User, error) {
	var user repository.User
	redisKey := fmt.Sprintf("user:%s", userID)
	err := r.Client.Get(ctx, redisKey).Scan(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// CreateRedisRepository ...
func CreateRedisRepository(url, password string) RedisService {
	return &redisService{
		Client: utilRedis.Initialize(&utilRedis.Config{
			Addr:     url,
			Password: password,
		}),
	}
}
