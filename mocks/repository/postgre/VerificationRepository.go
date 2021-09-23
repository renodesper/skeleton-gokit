// Code generated by mockery 2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	repository "gitlab.com/renodesper/gokit-microservices/repository"

	uuid "github.com/google/uuid"
)

// VerificationRepository is an autogenerated mock type for the VerificationRepository type
type VerificationRepository struct {
	mock.Mock
}

// CreateVerification provides a mock function with given fields: ctx, verificationPayload
func (_m *VerificationRepository) CreateVerification(ctx context.Context, verificationPayload *repository.Verification) (*repository.Verification, error) {
	ret := _m.Called(ctx, verificationPayload)

	var r0 *repository.Verification
	if rf, ok := ret.Get(0).(func(context.Context, *repository.Verification) *repository.Verification); ok {
		r0 = rf(ctx, verificationPayload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Verification)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *repository.Verification) error); ok {
		r1 = rf(ctx, verificationPayload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVerification provides a mock function with given fields: ctx, verificationType, token, isActive
func (_m *VerificationRepository) GetVerification(ctx context.Context, verificationType string, token string, isActive bool) (*repository.Verification, error) {
	ret := _m.Called(ctx, verificationType, token, isActive)

	var r0 *repository.Verification
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) *repository.Verification); ok {
		r0 = rf(ctx, verificationType, token, isActive)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Verification)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, bool) error); ok {
		r1 = rf(ctx, verificationType, token, isActive)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Invalidate provides a mock function with given fields: ctx, verificationID
func (_m *VerificationRepository) Invalidate(ctx context.Context, verificationID uuid.UUID) error {
	ret := _m.Called(ctx, verificationID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, verificationID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
