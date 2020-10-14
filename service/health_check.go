package service

import "github.com/spf13/viper"

// HealthCheck ...
func (us *UserSvc) HealthCheck() string {
	version := viper.GetString("app.version")
	if version == "" {
		return "1.0"
	}

	return version
}
