package repository

import (
	"time"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
)

// User ...
type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	IsActive  bool
	IsDeleted bool
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// MarshalBinary ...
func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

// UnmarshalBinary ...
func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
