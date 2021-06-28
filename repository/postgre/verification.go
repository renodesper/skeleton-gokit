package postgre

import (
	"context"

	"github.com/go-pg/pg/v10"
	"gitlab.com/renodesper/gokit-microservices/repository"
	"gitlab.com/renodesper/gokit-microservices/util/errors"
)

type (
	VerificationRepository interface {
		CreateVerification(ctx context.Context, verificationPayload *repository.Verification) (*repository.Verification, error)
	}

	VerificationRepo struct {
		Db *pg.DB
	}
)

func CreateVerificationRepository(db *pg.DB) VerificationRepository {
	return &VerificationRepo{
		Db: db,
	}
}

func (vr *VerificationRepo) CreateVerification(ctx context.Context, verificationPayload *repository.Verification) (*repository.Verification, error) {
	var verification repository.Verification

	_, err := vr.Db.WithContext(ctx).Model(verificationPayload).Returning("*").Insert(&verification)
	if err != nil {
		return nil, errors.FailedVerificationCreate.AppendError(err)
	}

	return &verification, nil
}
