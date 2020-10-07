package noop

import (
	"gitlab.com/renodesper/gokit-microservices/util/logger"
)

// CreateLogger creates logger that does nothing
func CreateLogger() logger.Logger {
	return logger.New()
}
