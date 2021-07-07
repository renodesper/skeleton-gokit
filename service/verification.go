package service

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"gitlab.com/renodesper/gokit-microservices/repository/postgre"
	"gitlab.com/renodesper/gokit-microservices/util/errors"
	"gitlab.com/renodesper/gokit-microservices/util/logger"
)

type (
	VerificationService interface {
		VerifyRegistration(ctx context.Context, token string) error
		VerifyResetPassword(ctx context.Context, userID uuid.UUID, token string) error
	}

	VerificationSvc struct {
		Log          logger.Logger
		Verification postgre.VerificationRepository
		EmailSvc     *EmailSvc
	}
)

// NewVerificationService creates auth service
func NewVerificationService(log logger.Logger, db *pg.DB) VerificationService {
	verificationRepo := postgre.CreateVerificationRepository(db)
	emailSvc := NewEmailSvc(log)

	return &VerificationSvc{
		Log:          log,
		Verification: verificationRepo,
		EmailSvc:     emailSvc,
	}
}

func (v *VerificationSvc) VerifyRegistration(ctx context.Context, token string) error {
	verification, err := v.Verification.GetVerification(ctx, "registration", token, true)
	if err != nil {
		return err
	}

	if verification.ExpiredAt.Before(time.Now()) {
		return errors.FailedVerificationExpiredToken
	}

	if verification.Token != token {
		return errors.FailedVerificationMismatchToken
	}

	verificationPayload := map[string]interface{}{
		"is_active":  false,
		"updated_at": time.Now(),
	}
	err = v.Verification.Invalidate(ctx, verification.ID, verificationPayload)
	if err != nil {
		return err
	}

	return nil
}

func (v *VerificationSvc) VerifyResetPassword(ctx context.Context, userID uuid.UUID, token string) error {
	return nil
}
