package service

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"gitlab.com/renodesper/gokit-microservices/repository"
	"gitlab.com/renodesper/gokit-microservices/repository/postgre"
	authUtil "gitlab.com/renodesper/gokit-microservices/util/auth"
	"gitlab.com/renodesper/gokit-microservices/util/constant"
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
		RequestResetPassword(ctx context.Context, email string) (*repository.User, error)
	}

	OauthSvc struct {
		Log          logger.Logger
		User         postgre.UserRepository
		Verification postgre.VerificationRepository
		EmailSvc     *EmailSvc
	}
)

// NewOauthService creates auth service
func NewOauthService(log logger.Logger, db *pg.DB) OauthService {
	userRepo := postgre.CreateUserRepository(log, db)
	verificationRepo := postgre.CreateVerificationRepository(log, db)
	emailSvc := NewEmailSvc(log)

	return &OauthSvc{
		Log:          log,
		User:         userRepo,
		Verification: verificationRepo,
		EmailSvc:     emailSvc,
	}
}

// NewOauthSvc creates auth service
func NewOauthSvc(log logger.Logger, db *pg.DB) *OauthSvc {
	userRepo := postgre.CreateUserRepository(log, db)
	emailSvc := NewEmailSvc(log)

	return &OauthSvc{
		Log:      log,
		User:     userRepo,
		EmailSvc: emailSvc,
	}
}

func (o *OauthSvc) Login(ctx context.Context, email string, password string) (*Token, error) {
	user, err := o.User.GetUserByEmail(ctx, email, repository.UserOptions{})
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
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
	userByEmail, err := o.User.GetUserByEmail(ctx, email, repository.UserOptions{})

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

	verificationToken := uuid.New()
	verificationPayload := repository.Verification{
		UserID:    user.ID,
		Type:      constant.VerificationTypeRegistration,
		Token:     verificationToken.String(),
		IsActive:  true,
		ExpiredAt: time.Now().Add(24 * time.Hour), // NOTE: Token is active for 1 day
	}
	_, err = o.Verification.CreateVerification(ctx, &verificationPayload)
	if err != nil {
		return nil, err
	}

	hEmail := o.EmailSvc.Welcome(user.Username, verificationToken.String())
	err = o.EmailSvc.SendMail(user.ID.String(), user.Email, constant.EmailSubjectWelcome, hEmail, constant.EmailTypeWelcome)
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (o *OauthSvc) RequestResetPassword(ctx context.Context, email string) (*repository.User, error) {
	// NOTE: Check for existing user
	user, err := o.User.GetUserByEmail(ctx, email, repository.UserOptions{})

	if err != nil {
		if wErr, ok := err.(e.Error); ok && wErr.Code != errors.FailedNoRows.Code {
			return nil, err
		}
	}

	verificationToken := uuid.New()
	verificationPayload := repository.Verification{
		UserID:    user.ID,
		Type:      constant.VerificationTypeResetPassword,
		Token:     verificationToken.String(),
		IsActive:  true,
		ExpiredAt: time.Now().Add(24 * time.Hour), // NOTE: Token is active for 1 day
	}
	_, err = o.Verification.CreateVerification(ctx, &verificationPayload)
	if err != nil {
		return nil, err
	}

	hEmail := o.EmailSvc.ResetPassword(user.Username, verificationToken.String())
	err = o.EmailSvc.SendMail(user.ID.String(), user.Email, constant.EmailSubjectResetPassword, hEmail, constant.EmailTypeResetPassword)
	if err != nil {
		return nil, err
	}

	return user, nil
}
