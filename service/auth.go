package service

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"gitlab.com/renodesper/gokit-microservices/repository"
	"gitlab.com/renodesper/gokit-microservices/repository/postgre"
	authUtil "gitlab.com/renodesper/gokit-microservices/util/auth"
	e "gitlab.com/renodesper/gokit-microservices/util/error"
	"gitlab.com/renodesper/gokit-microservices/util/errors"
	"gitlab.com/renodesper/gokit-microservices/util/logger"
	"golang.org/x/crypto/bcrypt"
)

type (
	OauthService interface {
		Login(ctx context.Context, email string, password string) (*Token, error)
		Logout(ctx context.Context, userID uuid.UUID) error
		Register(ctx context.Context, username string, email string, passwd string, isActive bool, isDeleted bool, isAdmin bool, createdFrom string) (*Token, error)
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

func (o *OauthSvc) Login(ctx context.Context, email string, password string) (*Token, error) {
	user, err := o.User.GetUserByEmail(ctx, email, repository.UserOptions{})
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
	if err != nil {
		o.Log.Info(err)
		return nil, errors.InvalidLoginCredential
	}

	token, err := authUtil.Token(user.ID)
	if err != nil {
		return nil, err
	}

	_, err = o.User.SetAccessToken(ctx, user.ID, token.AccessToken, token.RefreshToken, token.Expiry)
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
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

func (o *OauthSvc) Register(ctx context.Context, username string, email string, password string, isActive bool, isDeleted bool, isAdmin bool, createdFrom string) (*Token, error) {
	// NOTE: Check for existing username
	userByUsername, err := o.User.GetUserByUsername(ctx, username, repository.UserOptions{})

	if err != nil {
		if wErr, ok := err.(e.Error); ok && wErr.Code != errors.FailedNoRows.Code {
			return nil, err
		}
	}
	if userByUsername != nil {
		return nil, errors.FailedUsernameExist
	}

	// NOTE: Check for existing email
	userByEmail, _ := o.User.GetUserByEmail(ctx, email, repository.UserOptions{})

	if err != nil {
		if wErr, ok := err.(e.Error); ok && wErr.Code != errors.FailedNoRows.Code {
			return nil, err
		}
	}
	if userByEmail != nil {
		return nil, errors.FailedEmailExist
	}

	ID := uuid.New()

	passwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if createdFrom == "" {
		createdFrom = "Registration"
	}

	userPayload := repository.User{
		ID:          ID,
		Username:    username,
		Email:       email,
		Password:    string(passwd),
		IsActive:    isActive,
		IsDeleted:   isDeleted,
		IsAdmin:     isAdmin,
		CreatedFrom: createdFrom,
	}
	user, err := o.User.CreateUser(ctx, &userPayload)
	if err != nil {
		return nil, err
	}

	token, err := authUtil.Token(user.ID)
	if err != nil {
		return nil, err
	}

	_, err = o.User.SetAccessToken(ctx, user.ID, token.AccessToken, token.RefreshToken, token.Expiry)
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}
