package postgre

import (
	"context"
	"fmt"

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
		GetUser(ctx context.Context, userID uuid.UUID) (*repository.User, error)
		CreateUser(ctx context.Context, user *repository.User) (uuid.UUID, error)
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
		return nil, err
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

	q := ur.Db.Model(&users)

	if sort == "DESC" {
		if direction == "next" {
			q.Where("created_at <= ?", createdAt).Where("id < ?", userID).Order("created_at DESC")
		} else {
			q.Where("created_at >= ?", createdAt).Where("id > ?", userID).Order("created_at ASC")
		}
	} else {
		if direction == "next" {
			q.Where("created_at >= ?", createdAt).Where("id > ?", userID).Order("created_at ASC")
		} else {
			q.Where("created_at <= ?", createdAt).Where("id < ?", userID).Order("created_at DESC")
		}
	}

	err = q.Limit(limit).Select()
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

// GetUser ...
func (ur *UserRepo) GetUser(ctx context.Context, userID uuid.UUID) (*repository.User, error) {
	user := repository.User{}
	err := ur.Db.Model(&user).Where("id = ?", userID).Select()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser ...
func (ur *UserRepo) CreateUser(ctx context.Context, user *repository.User) (uuid.UUID, error) {
	var userID uuid.UUID

	_, err := ur.Db.WithContext(ctx).Model(user).Returning("id").Insert(&userID)
	if err != nil {
		return uuid.UUID{}, errors.FailedUserCreate.AppendError(err)
	}

	return userID, nil
}
