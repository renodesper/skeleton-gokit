package postgre

import (
	"context"
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"gitlab.com/renodesper/gokit-microservices/repository"
	"gitlab.com/renodesper/gokit-microservices/util/cursor"
	"gitlab.com/renodesper/gokit-microservices/util/errors"
)

type (
	UserRepository interface {
		GetAllUsers(ctx context.Context, sortBy string, sort string, skip int, limit int) ([]repository.User, error)
		GetAllUsersByCursor(ctx context.Context, cursor string, direction string, limit int, sort string) ([]repository.User, error)
		GetUserByID(ctx context.Context, userID uuid.UUID, opts repository.UserOptions) (*repository.User, error)
		GetUserByEmail(ctx context.Context, email string, opts repository.UserOptions) (*repository.User, error)
		GetUserByUsername(ctx context.Context, username string, opts repository.UserOptions) (*repository.User, error)
		CreateUser(ctx context.Context, userPayload *repository.User) (*repository.User, error)
		UpdateUser(ctx context.Context, userID uuid.UUID, userPayload map[string]interface{}) (*repository.User, error)
		SetAccessToken(ctx context.Context, userID uuid.UUID, accessToken string, refreshToken string) (*repository.User, error)
		SetUserStatus(ctx context.Context, userID uuid.UUID, isActive bool) (*repository.User, error)
		SetUserRole(ctx context.Context, userID uuid.UUID, isAdmin bool) (*repository.User, error)
		SetUserExpiry(ctx context.Context, userID uuid.UUID, expiredAt time.Time) (*repository.User, error)
		DeleteUser(ctx context.Context, userID uuid.UUID) (*repository.User, error)
	}

	UserRepo struct {
		Db *pg.DB
	}
)

// CreateUserRepository creates user repository
func CreateUserRepository(db *pg.DB) UserRepository {
	return &UserRepo{
		Db: db,
	}
}

// GetAllUsers ...
func (ur *UserRepo) GetAllUsers(ctx context.Context, sortBy string, sort string, skip int, limit int) ([]repository.User, error) {
	var users []repository.User

	if sortBy == "" {
		sortBy = "created_at"
	}
	if sort == "" {
		sort = "DESC"
	}
	order := fmt.Sprintf("%s %s", sortBy, sort)

	err := ur.Db.Model(&users).Limit(limit).Offset(skip).Order(order).Select()
	if err != nil {
		return nil, errors.FailedUsersFetch.AppendError(err)
	}

	return users, nil
}

// GetAllUsersByCursor will be more effective for feed type list
func (ur *UserRepo) GetAllUsersByCursor(ctx context.Context, sort string, direction string, limit int, encodedCursor string) ([]repository.User, error) {
	var users []repository.User

	createdAt, userIDStr, err := cursor.DecodeCursor(encodedCursor)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, err
	}

	sql := ur.Db.Model(&users)

	if sort == "DESC" {
		if direction == "next" {
			sql.Where("created_at <= ?", createdAt).Where("id < ?", userID).Order("created_at DESC")
		} else {
			sql.Where("created_at >= ?", createdAt).Where("id > ?", userID).Order("created_at ASC")
		}
	} else {
		if direction == "next" {
			sql.Where("created_at >= ?", createdAt).Where("id > ?", userID).Order("created_at ASC")
		} else {
			sql.Where("created_at <= ?", createdAt).Where("id < ?", userID).Order("created_at DESC")
		}
	}

	err = sql.Limit(limit).Select()
	if err != nil {
		return nil, err
	}

	if direction != "next" {
		newUsers := make([]repository.User, 0, len(users))
		for i := len(users) - 1; i >= 0; i-- {
			newUsers = append(newUsers, users[i])
		}
		return newUsers, nil
	}

	return users, nil
}

// GetUserByID ...
func (ur *UserRepo) GetUserByID(ctx context.Context, userID uuid.UUID, opts repository.UserOptions) (*repository.User, error) {
	user := repository.User{}

	sql := ur.Db.Model(&user).Where("id = ?", userID)
	if opts.IsActive != nil {
		sql.Where("is_active = ?", *opts.IsActive)
	}
	if opts.IsDeleted != nil {
		sql.Where("is_deleted = ?", *opts.IsDeleted)
	}
	if opts.IsAdmin != nil {
		sql.Where("is_admin = ?", *opts.IsAdmin)
	}
	if opts.CreatedFrom != nil {
		sql.Where("created_from = ?", *opts.CreatedFrom)
	}

	err := sql.Select()
	if err != nil {
		return nil, errors.FailedUserFetch.AppendError(err)
	}

	return &user, nil
}

// GetUserByEmail ...
func (ur *UserRepo) GetUserByEmail(ctx context.Context, email string, opts repository.UserOptions) (*repository.User, error) {
	user := repository.User{}

	sql := ur.Db.Model(&user).Where("email = ?", email)
	if opts.IsActive != nil {
		sql.Where("is_active = ?", *opts.IsActive)
	}
	if opts.IsDeleted != nil {
		sql.Where("is_deleted = ?", *opts.IsDeleted)
	}
	if opts.IsAdmin != nil {
		sql.Where("is_admin = ?", *opts.IsAdmin)
	}
	if opts.CreatedFrom != nil {
		sql.Where("created_from = ?", *opts.CreatedFrom)
	}

	err := sql.Select()
	if err != nil {
		return nil, errors.FailedUserFetch.AppendError(err)
	}

	return &user, nil
}

// GetUserByUsername ...
func (ur *UserRepo) GetUserByUsername(ctx context.Context, username string, opts repository.UserOptions) (*repository.User, error) {
	user := repository.User{}

	sql := ur.Db.Model(&user).Where("username = ?", username)
	if opts.IsActive != nil {
		sql.Where("is_active = ?", *opts.IsActive)
	}
	if opts.IsDeleted != nil {
		sql.Where("is_deleted = ?", *opts.IsDeleted)
	}
	if opts.IsAdmin != nil {
		sql.Where("is_admin = ?", *opts.IsAdmin)
	}
	if opts.CreatedFrom != nil {
		sql.Where("created_from = ?", *opts.CreatedFrom)
	}

	err := sql.Select()
	if err != nil {
		return nil, errors.FailedUserFetch.AppendError(err)
	}

	return &user, nil
}

// CreateUser ...
func (ur *UserRepo) CreateUser(ctx context.Context, userPayload *repository.User) (*repository.User, error) {
	var user repository.User

	_, err := ur.Db.WithContext(ctx).Model(userPayload).Returning("*").Insert(&user)
	if err != nil {
		return nil, errors.FailedUserCreate.AppendError(err)
	}

	return &user, nil
}

func (ur *UserRepo) UpdateUser(ctx context.Context, userID uuid.UUID, userPayload map[string]interface{}) (*repository.User, error) {
	var user repository.User
	_, err := ur.Db.Model(&userPayload).TableExpr("users").Where("id = ?", userID).Returning("*").Update(&user)
	if err != nil {
		return nil, errors.FailedUserUpdate.AppendError(err)
	}

	return &user, nil
}

func (ur *UserRepo) SetAccessToken(ctx context.Context, userID uuid.UUID, accessToken string, refreshToken string) (*repository.User, error) {
	userPayload := map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	var user repository.User
	_, err := ur.Db.Model(&userPayload).TableExpr("users").Where("id = ?", userID).Returning("*").Update(&user)
	if err != nil {
		return nil, errors.FailedUserUpdate.AppendError(err)
	}

	return &user, nil
}

func (ur *UserRepo) SetUserStatus(ctx context.Context, userID uuid.UUID, isActive bool) (*repository.User, error) {
	userPayload := map[string]interface{}{
		"is_active": isActive,
	}

	var user repository.User
	_, err := ur.Db.Model(&userPayload).TableExpr("users").Where("id = ?", userID).Returning("*").Update(&user)
	if err != nil {
		return nil, errors.FailedUserUpdate.AppendError(err)
	}

	return &user, nil
}

func (ur *UserRepo) SetUserRole(ctx context.Context, userID uuid.UUID, isAdmin bool) (*repository.User, error) {
	userPayload := map[string]interface{}{
		"is_admin": isAdmin,
	}

	var user repository.User
	_, err := ur.Db.Model(&userPayload).TableExpr("users").Where("id = ?", userID).Returning("*").Update(&user)
	if err != nil {
		return nil, errors.FailedUserUpdate.AppendError(err)
	}

	return &user, nil
}

func (ur *UserRepo) SetUserExpiry(ctx context.Context, userID uuid.UUID, expiredAt time.Time) (*repository.User, error) {
	userPayload := map[string]interface{}{
		"expired_at": expiredAt,
	}

	var user repository.User
	_, err := ur.Db.Model(&userPayload).TableExpr("users").Where("id = ?", userID).Returning("*").Update(&user)
	if err != nil {
		return nil, errors.FailedUserUpdate.AppendError(err)
	}

	return &user, nil
}

func (ur *UserRepo) DeleteUser(ctx context.Context, userID uuid.UUID) (*repository.User, error) {
	userPayload := map[string]interface{}{
		"is_deleted": true,
	}

	var user repository.User
	_, err := ur.Db.Model(&userPayload).TableExpr("users").Where("id = ?", userID).Returning("*").Update(&user)
	if err != nil {
		return nil, errors.FailedUserDelete.AppendError(err)
	}

	return &user, nil
}
