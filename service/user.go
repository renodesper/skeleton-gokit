package service

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"gitlab.com/renodesper/gokit-microservices/repository"
	"gitlab.com/renodesper/gokit-microservices/repository/postgre"
	"golang.org/x/crypto/bcrypt"
)

type (
	// UserService ...
	UserService interface {
		GetAllUsers(ctx context.Context, sortBy string, sort string, skip int, limit int) ([]repository.User, error)
		GetUser(ctx context.Context, userIDStr string) (*repository.User, error)
		CreateUser(ctx context.Context, req *CreateUserRequest) (uuid.UUID, error)
	}

	UserSvc struct {
		User postgre.UserRepository
	}

	CreateUserRequest struct {
		Username string
		Email    string
		Password string
		IsAdmin  bool
	}

	UpdateUserRequest struct {
	}
)

// NewUserService creates user service
func NewUserService(db *pg.DB) UserService {
	userRepo := postgre.CreateUserRepository(db)
	return &UserSvc{
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

func (us *UserSvc) GetUser(ctx context.Context, userIDStr string) (*repository.User, error) {
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, err
	}

	user, err := us.User.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserSvc) CreateUser(ctx context.Context, req *CreateUserRequest) (uuid.UUID, error) {
	ID := uuid.New()

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.UUID{}, err
	}

	userPayload := repository.User{
		ID:        ID,
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(password),
		IsActive:  false,
		IsDeleted: false,
		IsAdmin:   req.IsAdmin,
	}
	userID, err := us.User.CreateUser(ctx, &userPayload)
	if err != nil {
		return userID, err
	}

	return userID, nil
}

func (us *UserSvc) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*repository.User, error) {
	return nil, nil
}

func (us *UserSvc) DeleteUser(ctx context.Context, userID string) (*repository.User, error) {
	return nil, nil
}
