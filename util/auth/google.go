package auth

import (
	"fmt"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var config *oauth2.Config

func GetGoogleOauthConfig() *oauth2.Config {
	if config != nil {
		return config
	}

	scheme := viper.GetString("app.scheme")
	host := viper.GetString("app.host")
	port := viper.GetInt("app.port")
	redirectURL := fmt.Sprintf("%s://%s:%d/auth/google/callback", scheme, host, port)

	config = &oauth2.Config{
		ClientID:     viper.GetString("google.client_id"),
		ClientSecret: viper.GetString("google.client_secret"),
		RedirectURL:  redirectURL,
		Scopes:       []string{"email"}, // NOTE: https://www.googleapis.com/auth/userinfo.email
		Endpoint:     google.Endpoint,
	}

	return config
}
