package service

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"gitlab.com/renodesper/gokit-microservices/repository"
	"gitlab.com/renodesper/gokit-microservices/repository/postgre"
	authUtil "gitlab.com/renodesper/gokit-microservices/util/auth"
	"gitlab.com/renodesper/gokit-microservices/util/errors"
	"gitlab.com/renodesper/gokit-microservices/util/logger"
)

type (
	GoogleOauthService interface {
		OauthCallback(ctx context.Context, code string) (*Token, error)
		GetUserData(ctx context.Context, code string) (*GoogleUser, error)
	}

	GoogleOauthSvc struct {
		Log      logger.Logger
		User     postgre.UserRepository
		OauthSvc *OauthSvc
	}

	GoogleUser struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Picture       string `json:"picture"`
	}

	Token struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

// NewGoogleAuthService creates auth google service
func NewGoogleOauthService(log logger.Logger, db *pg.DB) GoogleOauthService {
	userRepo := postgre.CreateUserRepository(db)
	oauthSvc := NewOauthSvc(log, db)

	return &GoogleOauthSvc{
		Log:      log,
		User:     userRepo,
		OauthSvc: oauthSvc,
	}
}

func (g *GoogleOauthSvc) OauthCallback(ctx context.Context, code string) (*Token, error) {
	googleUser, err := g.GetUserData(ctx, code)
	if err != nil {
		return nil, err
	}

	user, _ := g.User.GetUserByEmail(ctx, googleUser.Email, repository.UserOptions{})
	if user == nil {
		s := strings.Split(googleUser.Email, "@")
		username := s[0]

		password := uuid.New().String()

		token, err := g.OauthSvc.Register(ctx, username, googleUser.Email, password, false, false, false, "GoogleOauth")
		if err != nil {
			return nil, err
		}

		return token, nil
	}

	token, err := authUtil.Token(user.ID)
	if err != nil {
		return nil, err
	}

	_, err = g.User.SetAccessToken(ctx, user.ID, token.AccessToken, token.RefreshToken, token.Expiry)
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (g *GoogleOauthSvc) GetUserData(ctx context.Context, code string) (*GoogleUser, error) {
	config := authUtil.GetGoogleOauthConfig()

	token, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, errors.InvalidGoogleOauthCodeExchange
	}

	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, errors.FailedGoogleUserFetch
	}

	var googleUser GoogleUser
	if err := json.NewDecoder(response.Body).Decode(&googleUser); err != nil {
		return nil, errors.UnparsableJSON
	}
	defer response.Body.Close()

	return &googleUser, nil
}
