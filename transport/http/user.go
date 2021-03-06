package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/go-zoo/bone"
	"github.com/google/uuid"
	"github.com/iancoleman/strcase"
	"gitlab.com/renodesper/gokit-microservices/endpoint"
	"gitlab.com/renodesper/gokit-microservices/util/errors"
)

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.UnparsableJSON
	}
	defer r.Body.Close()

	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, errors.InvalidRequest
	}

	return req, nil
}

func decodeGetAllUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	sortBy := "created_at"
	sort := "DESC"
	skip := 0
	limit := 10

	var req endpoint.GetAllUsersRequest

	sortByParam := r.URL.Query().Get("sortBy")
	if sortByParam != "" {
		sortBy = strcase.ToSnake(sortByParam)
	}
	req.SortBy = sortBy

	sortParam := r.URL.Query().Get("sort")
	if sortParam != "" {
		sort = sortParam
	}
	req.Sort = sort

	skipParam := r.URL.Query().Get("skip")
	if skipParam != "" {
		skip, _ = strconv.Atoi(skipParam)
	}
	req.Skip = skip

	limitParam := r.URL.Query().Get("limit")
	if limitParam != "" {
		limit, _ = strconv.Atoi(limitParam)
	}
	req.Limit = limit

	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, errors.InvalidRequest
	}

	return req, nil
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	IDStr := bone.GetValue(r, "id")

	var req endpoint.GetUserRequest

	ID, err := uuid.Parse(IDStr)
	if err != nil {
		return nil, err
	}

	req.ID = ID

	return req, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.UpdateUserRequest

	IDStr := bone.GetValue(r, "id")
	ID, err := uuid.Parse(IDStr)
	if err != nil {
		return nil, err
	}

	req.ID = ID

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.UnparsableJSON
	}
	defer r.Body.Close()

	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, errors.InvalidRequest
	}

	return req, nil
}

func decodeSetPasswordRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.SetPasswordRequest

	IDStr := bone.GetValue(r, "id")
	ID, err := uuid.Parse(IDStr)
	if err != nil {
		return nil, err
	}

	req.ID = ID

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.UnparsableJSON
	}
	defer r.Body.Close()

	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, errors.InvalidRequest
	}

	return req, nil
}

func decodeSetAccessTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.SetAccessTokenRequest

	IDStr := bone.GetValue(r, "id")
	ID, err := uuid.Parse(IDStr)
	if err != nil {
		return nil, err
	}

	req.ID = ID

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.UnparsableJSON
	}
	defer r.Body.Close()

	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, errors.InvalidRequest
	}

	return req, nil
}

func decodeSetUserStatusRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.SetUserStatusRequest

	IDStr := bone.GetValue(r, "id")
	ID, err := uuid.Parse(IDStr)
	if err != nil {
		return nil, err
	}

	req.ID = ID

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.UnparsableJSON
	}
	defer r.Body.Close()

	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, errors.InvalidRequest
	}

	return req, nil
}

func decodeSetUserRoleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.SetUserRoleRequest

	IDStr := bone.GetValue(r, "id")
	ID, err := uuid.Parse(IDStr)
	if err != nil {
		return nil, err
	}

	req.ID = ID

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.UnparsableJSON
	}
	defer r.Body.Close()

	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, errors.InvalidRequest
	}

	return req, nil
}

func decodeSetUserExpiryEndpoint(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.SetUserExpiryRequest

	IDStr := bone.GetValue(r, "id")
	ID, err := uuid.Parse(IDStr)
	if err != nil {
		return nil, err
	}

	req.ID = ID

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.UnparsableJSON
	}
	defer r.Body.Close()

	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, errors.InvalidRequest
	}

	return req, nil
}

func decodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.GetUserRequest

	IDStr := bone.GetValue(r, "id")
	ID, err := uuid.Parse(IDStr)
	if err != nil {
		return nil, err
	}

	req.ID = ID

	return req, nil
}
