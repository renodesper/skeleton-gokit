package service

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"gitlab.com/renodesper/gokit-microservices/repository"
	"gitlab.com/renodesper/gokit-microservices/repository/postgre"
	"gitlab.com/renodesper/gokit-microservices/util/constant"
	"gitlab.com/renodesper/gokit-microservices/util/errors"
	"gitlab.com/renodesper/gokit-microservices/util/logger"
)

type (
	VerificationService interface {
		VerifyRegistration(ctx context.Context, token string) (uuid.UUID, error)
		VerifyResetPassword(ctx context.Context, token, password, verifyPassword string) (uuid.UUID, error)
		VerifyToken(ctx context.Context, token string, verification *repository.Verification) (uuid.UUID, error)
	}

	VerificationSvc struct {
		Log          logger.Logger
		Verification postgre.VerificationRepository
		EmailSvc     EmailService
		UserSvc      UserService
	}
)

// NewVerificationService creates auth service
func NewVerificationService(log logger.Logger, db *pg.DB) VerificationService {
	verificationRepo := postgre.CreateVerificationRepository(log, db)
	emailSvc := NewEmailService(log)
	userSvc := NewUserService(log, db)

	return &VerificationSvc{
		Log:          log,
		Verification: verificationRepo,
		EmailSvc:     emailSvc,
		UserSvc:      userSvc,
	}
}

func (v *VerificationSvc) VerifyRegistration(ctx context.Context, token string) (uuid.UUID, error) {
	isActive := true
	verification, err := v.Verification.GetVerification(ctx, constant.VerificationTypeRegistration, token, isActive)
	if err != nil {
		return uuid.Nil, err
	}

	userID, err := v.VerifyToken(ctx, token, verification)
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}

func (v *VerificationSvc) VerifyResetPassword(ctx context.Context, token, password, verifyPassword string) (uuid.UUID, error) {
	isActive := true
	verification, err := v.Verification.GetVerification(ctx, constant.VerificationTypeResetPassword, token, isActive)
	if err != nil {
		return uuid.Nil, err
	}

	userID, err := v.VerifyToken(ctx, token, verification)
	if err != nil {
		return uuid.Nil, err
	}

	user, err := v.UserSvc.SetPassword(ctx, userID, password, verifyPassword)
	if err != nil {
		return uuid.Nil, err
	}

	err = v.Verification.Invalidate(ctx, verification.ID)
	if err != nil {
		return uuid.Nil, err
	}

	hEmail := v.EmailSvc.ResetPasswordNotification(user.Username)
	err = v.EmailSvc.SendMail(
		user.ID.String(),
		user.Email,
		constant.EmailSubjectResetPasswordNotification,
		hEmail,
		constant.EmailTypeResetPasswordNotification,
	)
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}

func (v *VerificationSvc) VerifyToken(ctx context.Context, token string, verification *repository.Verification) (uuid.UUID, error) {
	if verification.ExpiredAt.Before(time.Now()) {
		return uuid.Nil, errors.FailedVerificationExpiredToken
	}

	if verification.Token != token {
		return uuid.Nil, errors.FailedVerificationMismatchToken
	}

	err := v.Verification.Invalidate(ctx, verification.ID)
	if err != nil {
		return uuid.Nil, err
	}

	return verification.UserID, nil
}
