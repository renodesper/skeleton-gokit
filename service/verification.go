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
		VerifyResetPassword(ctx context.Context, token string) (uuid.UUID, error)
		VerifyToken(ctx context.Context, token string, verification *repository.Verification) (uuid.UUID, error)
	}

	VerificationSvc struct {
		Log          logger.Logger
		Verification postgre.VerificationRepository
		EmailSvc     *EmailSvc
	}
)

// NewVerificationService creates auth service
func NewVerificationService(log logger.Logger, db *pg.DB) VerificationService {
	verificationRepo := postgre.CreateVerificationRepository(log, db)
	emailSvc := NewEmailSvc(log)

	return &VerificationSvc{
		Log:          log,
		Verification: verificationRepo,
		EmailSvc:     emailSvc,
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

func (v *VerificationSvc) VerifyResetPassword(ctx context.Context, token string) (uuid.UUID, error) {
	isActive := true
	verification, err := v.Verification.GetVerification(ctx, constant.VerificationTypeResetPassword, token, isActive)
	if err != nil {
		return uuid.Nil, err
	}

	userID, err := v.VerifyToken(ctx, token, verification)
	if err != nil {
		return uuid.Nil, err
	}

	verificationToken := uuid.New()
	verificationPayload := repository.Verification{
		UserID:    userID,
		Type:      constant.VerificationTypeUpdatePassword,
		Token:     verificationToken.String(),
		IsActive:  true,
		ExpiredAt: time.Now().Add(24 * time.Hour), // NOTE: Token is active for 1 day
	}
	_, err = v.Verification.CreateVerification(ctx, &verificationPayload)
	if err != nil {
		return uuid.Nil, err
	}

	return verificationToken, nil
}

func (v *VerificationSvc) VerifyToken(ctx context.Context, token string, verification *repository.Verification) (uuid.UUID, error) {
	if verification.ExpiredAt.Before(time.Now()) {
		return uuid.Nil, errors.FailedVerificationExpiredToken
	}

	if verification.Token != token {
		return uuid.Nil, errors.FailedVerificationMismatchToken
	}

	verificationPayload := map[string]interface{}{
		"is_active":  false,
		"updated_at": time.Now(),
	}
	err := v.Verification.Invalidate(ctx, verification.ID, verificationPayload)
	if err != nil {
		return uuid.Nil, err
	}

	return verification.UserID, nil
}
