package repository

import (
	"time"

	"github.com/google/uuid"
)

type (
	Verification struct {
		TableName struct{} `pg:"verification"`

		ID        uuid.UUID `db:"id" json:"id"`
		UserID    uuid.UUID `db:"userId" json:"userId"`
		Type      string    `db:"type" json:"type"`
		Token     string    `db:"token" json:"token"`
		IsActive  bool      `db:"isActive" json:"isActive"`
		ExpiredAt time.Time `db:"expiredAt" json:"expiredAt"`
		CreatedAt time.Time `db:"createdAt" json:"createdAt"`
		UpdatedAt time.Time `db:"updatedAt" json:"updatedAt"`
	}

	VerificationOptions struct {
		IsActive *bool
	}
)

// MarshalBinary ...
func (v *Verification) MarshalBinary() ([]byte, error) {
	return json.Marshal(v)
}

// UnmarshalBinary ...
func (v *Verification) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, v)
}
