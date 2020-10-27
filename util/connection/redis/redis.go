package redis

import (
	"time"

	"github.com/go-redis/redis/v8"
)

type (
	// Config ...
	Config struct {
		Addr        string
		Password    string
		DB          int
		IdleTimeout time.Duration
	}
)

// Initialize ...
func Initialize(c *Config) *redis.Client {
	if c.Addr == "" {
		c.Addr = "localhost:6379"
	}

	if c.IdleTimeout == 0 {
		c.IdleTimeout = 300 * time.Second
	}

	options := redis.Options{
		Addr:        c.Addr,
		Password:    c.Password,
		DB:          c.DB,
		IdleTimeout: c.IdleTimeout,
	}

	return redis.NewClient(&options)
}
