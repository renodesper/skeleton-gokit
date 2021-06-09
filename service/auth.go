package service

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"gitlab.com/renodesper/gokit-microservices/repository/postgre"
	authUtil "gitlab.com/renodesper/gokit-microservices/util/auth"
	"gitlab.com/renodesper/gokit-microservices/util/logger"
)

type (
	OauthService interface {
		Logout(ctx context.Context, userID uuid.UUID) error
	}

	OauthSvc struct {
		Log  logger.Logger
		User postgre.UserRepository
	}
)

// NewOauthService creates auth service
func NewOauthService(log logger.Logger, db *pg.DB) OauthService {
	userRepo := postgre.CreateUserRepository(db)
	return &OauthSvc{
		Log:  log,
		User: userRepo,
	}
}

func (o *OauthSvc) Logout(ctx context.Context, userID uuid.UUID) error {
	token, err := authUtil.Token(userID)
	if err != nil {
		return err
	}

	_, err = o.User.SetAccessToken(ctx, userID, "", token.RefreshToken, time.Now())
	if err != nil {
		return err
	}

	return nil
}
