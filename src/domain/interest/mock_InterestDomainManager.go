// Code generated by mockery v2.14.0. DO NOT EDIT.

package interest

import (
	context "context"

	entity "github.com/abisatria1/simple-dating-app/src/domain/entity"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
)

// MockInterestDomainManager is an autogenerated mock type for the InterestDomainManager type
type MockInterestDomainManager struct {
	mock.Mock
}

// Begin provides a mock function with given fields: ctx
func (_m *MockInterestDomainManager) Begin(ctx context.Context) *gorm.DB {
	ret := _m.Called(ctx)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(context.Context) *gorm.DB); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// GetAllInterests provides a mock function with given fields: ctx
func (_m *MockInterestDomainManager) GetAllInterests(ctx context.Context) ([]entity.Interest, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Interest
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Interest); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Interest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockInterestDomainManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockInterestDomainManager creates a new instance of MockInterestDomainManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockInterestDomainManager(t mockConstructorTestingTNewMockInterestDomainManager) *MockInterestDomainManager {
	mock := &MockInterestDomainManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}