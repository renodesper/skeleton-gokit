package repository

import (
	"time"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
)

type (
	User struct {
		ID           uuid.UUID `db:"id" json:"id"`
		Username     string    `db:"username" json:"username"`
		Email        string    `db:"email" validate:"email" json:"email" validate:"email"`
		Password     string    `db:"password" json:"password"`
		IsActive     bool      `db:"isActive" json:"isActive"`
		IsDeleted    bool      `db:"isDeleted" json:"isDeleted"`
		IsAdmin      bool      `db:"isAdmin" json:"isAdmin"`
		AccessToken  string    `db:"accessToken" json:"accessToken"`
		RefreshToken string    `db:"refreshToken" json:"refreshToken"`
		CreatedFrom  string    `db:"createdFrom" json:"createdFrom"`
		ExpiredAt    time.Time `db:"expiredAt" json:"expiredAt"`
		CreatedAt    time.Time `db:"createdAt" json:"createdAt"`
		UpdatedAt    time.Time `db:"updatedAt" json:"updatedAt"`
	}

	UserOptions struct {
		IsActive    *bool
		IsDeleted   *bool
		IsAdmin     *bool
		CreatedFrom *string
	}
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// MarshalBinary ...
func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

// UnmarshalBinary ...
func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
