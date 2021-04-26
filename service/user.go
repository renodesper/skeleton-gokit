package service

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"gitlab.com/renodesper/gokit-microservices/repository"
	"gitlab.com/renodesper/gokit-microservices/repository/postgre"
	"gitlab.com/renodesper/gokit-microservices/util/errors"
	"gitlab.com/renodesper/gokit-microservices/util/logger"
	"golang.org/x/crypto/bcrypt"
)

type (
	// UserService ...
	UserService interface {
		GetAllUsers(ctx context.Context, sortBy string, sort string, skip int, limit int) ([]repository.User, error)
		GetUser(ctx context.Context, userID uuid.UUID) (*repository.User, error)
		CreateUser(ctx context.Context, payload *CreateUserRequest) (*repository.User, error)
		UpdateUser(ctx context.Context, userID uuid.UUID, payload *UpdateUserRequest) (*repository.User, error)
		SetAccessToken(ctx context.Context, userID uuid.UUID, accessToken string, refreshToken string) (*repository.User, error)
		SetUserStatus(ctx context.Context, userID uuid.UUID, isActive bool) (*repository.User, error)
		SetUserRole(ctx context.Context, userID uuid.UUID, isAdmin bool) (*repository.User, error)
		SetUserExpiry(ctx context.Context, userID uuid.UUID, expiredAt time.Time) (*repository.User, error)
		DeleteUser(ctx context.Context, userID uuid.UUID) (*repository.User, error)
	}

	UserSvc struct {
		Log  logger.Logger
		User postgre.UserRepository
	}

	CreateUserRequest struct {
		Username string
		Email    string
		Password string
		IsAdmin  bool
	}

	UpdateUserRequest struct {
		Username string
		Email    string
		Password string
	}
)

// NewUserService creates user service
func NewUserService(log logger.Logger, db *pg.DB) UserService {
	userRepo := postgre.CreateUserRepository(db)
	return &UserSvc{
		Log:  log,
		User: userRepo,
	}
}

func (us *UserSvc) GetAllUsers(ctx context.Context, sortBy string, sort string, skip int, limit int) ([]repository.User, error) {
	var users []repository.User

	users, err := us.User.GetAllUsers(ctx, sortBy, sort, skip, limit)
	if err != nil {
		return nil, err
	}

	if users == nil {
		users = []repository.User{}
	}

	return users, nil
}

func (us *UserSvc) GetUser(ctx context.Context, userID uuid.UUID) (*repository.User, error) {
	user, err := us.User.GetUserByID(ctx, userID, repository.UserOptions{})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserSvc) CreateUser(ctx context.Context, payload *CreateUserRequest) (*repository.User, error) {
	if payload.Email != "" {
		user, _ := us.User.GetUserByEmail(ctx, payload.Email, repository.UserOptions{})

		if user != nil {
			return nil, errors.FailedEmailExist
		}
	}

	if payload.Username != "" {
		user, _ := us.User.GetUserByUsername(ctx, payload.Username, repository.UserOptions{})

		if user != nil {
			return nil, errors.FailedUsernameExist
		}
	}

	ID := uuid.New()

	password, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	userPayload := repository.User{
		ID:          ID,
		Username:    payload.Username,
		Email:       payload.Email,
		Password:    string(password),
		IsActive:    false,
		IsDeleted:   false,
		IsAdmin:     payload.IsAdmin,
		CreatedFrom: "UserAPI",
	}
	user, err := us.User.CreateUser(ctx, &userPayload)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserSvc) UpdateUser(ctx context.Context, userID uuid.UUID, payload *UpdateUserRequest) (*repository.User, error) {
	if payload.Email != "" {
		user, _ := us.User.GetUserByEmail(ctx, payload.Email, repository.UserOptions{})

		if user != nil && userID != user.ID {
			return nil, errors.FailedEmailExist
		}
	}

	if payload.Username != "" {
		user, _ := us.User.GetUserByUsername(ctx, payload.Username, repository.UserOptions{})

		if user != nil && userID != user.ID {
			return nil, errors.FailedUsernameExist
		}
	}

	var userPayload = make(map[string]interface{})

	if payload.Username != "" {
		userPayload["username"] = payload.Username
	}

	if payload.Email != "" {
		userPayload["email"] = payload.Email
	}

	if payload.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		userPayload["password"] = password
	}

	user, err := us.User.UpdateUser(ctx, userID, userPayload)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserSvc) SetAccessToken(ctx context.Context, userID uuid.UUID, accessToken string, refreshToken string) (*repository.User, error) {
	user, err := us.User.SetAccessToken(ctx, userID, accessToken, refreshToken)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserSvc) SetUserStatus(ctx context.Context, userID uuid.UUID, isActive bool) (*repository.User, error) {
	user, err := us.User.SetUserStatus(ctx, userID, isActive)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserSvc) SetUserRole(ctx context.Context, userID uuid.UUID, isAdmin bool) (*repository.User, error) {
	user, err := us.User.SetUserRole(ctx, userID, isAdmin)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserSvc) SetUserExpiry(ctx context.Context, userID uuid.UUID, expiredAt time.Time) (*repository.User, error) {
	user, err := us.User.SetUserExpiry(ctx, userID, expiredAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserSvc) DeleteUser(ctx context.Context, userID uuid.UUID) (*repository.User, error) {
	user, err := us.User.DeleteUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
