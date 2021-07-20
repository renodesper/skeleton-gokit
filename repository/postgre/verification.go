package postgre

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"gitlab.com/renodesper/gokit-microservices/repository"
	"gitlab.com/renodesper/gokit-microservices/util/errors"
	"gitlab.com/renodesper/gokit-microservices/util/logger"
)

type (
	VerificationRepository interface {
		GetVerification(ctx context.Context, verificationType, token string, isActive bool) (*repository.Verification, error)
		CreateVerification(ctx context.Context, verificationPayload *repository.Verification) (*repository.Verification, error)
		Invalidate(ctx context.Context, verificationID uuid.UUID, verificationPayload map[string]interface{}) error
	}

	VerificationRepo struct {
		Log logger.Logger
		Db  *pg.DB
	}
)

var verificationTable = "verification"

func CreateVerificationRepository(log logger.Logger, db *pg.DB) VerificationRepository {
	return &VerificationRepo{
		Log: log,
		Db:  db,
	}
}

// GetVerification ...
func (vr *VerificationRepo) GetVerification(ctx context.Context, verificationType, token string, isActive bool) (*repository.Verification, error) {
	verification := repository.Verification{}

	sql := vr.Db.WithContext(ctx).Model(&verification).
		Where("type = ?", verificationType).
		Where("token =?", token).
		Where("is_active = ?", isActive)

	err := sql.Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, errors.FailedNoRows.AppendError(err)
		}

		return nil, errors.FailedVerificationFetch.AppendError(err)
	}

	return &verification, nil
}

func (vr *VerificationRepo) CreateVerification(ctx context.Context, verificationPayload *repository.Verification) (*repository.Verification, error) {
	var verification repository.Verification

	b, _ := json.Marshal(verificationPayload)
	fmt.Println(string(b))

	_, err := vr.Db.WithContext(ctx).Model(verificationPayload).Returning("*").Insert(&verification)
	if err != nil {
		return nil, errors.FailedVerificationCreate.AppendError(err)
	}

	return &verification, nil
}

func (vr *VerificationRepo) Invalidate(ctx context.Context, verificationID uuid.UUID, verificationPayload map[string]interface{}) error {
	var verification repository.Verification
	_, err := vr.Db.WithContext(ctx).Model(&verificationPayload).Table(verificationTable).Where("id = ?", verificationID).Returning("*").Update(&verification)
	if err != nil {
		return errors.FailedVerificationUpdate.AppendError(err)
	}

	return nil
}
