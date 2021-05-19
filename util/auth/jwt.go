package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// JWTClaims struct
type JWTClaims struct {
	Identity string `json:"identity"`
	jwt.StandardClaims
}

var JWT_SECRET_KEY = viper.GetString("jwt.secret_key")

// ParseJWTWithClaims ...
func ParseJWTWithClaims(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET_KEY), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims.Identity, nil
	}

	return "", err
}

// Token ...
func Token(userID uuid.UUID) (*oauth2.Token, error) {
	iat := time.Now()
	exp := iat.Add(time.Hour * 24 * 365) // NOTE: 1 year
	claims := JWTClaims{
		userID.String(),
		jwt.StandardClaims{
			IssuedAt:  iat.Unix(),
			NotBefore: iat.Unix(),
			ExpiresAt: exp.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return nil, err
	}

	return &oauth2.Token{
		AccessToken:  tokenString,
		RefreshToken: uuid.New().String(),
		TokenType:    "Bearer",
		Expiry:       exp,
	}, nil
}
