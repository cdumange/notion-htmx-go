// Code generated by mockery v2.33.3. DO NOT EDIT.

package routing

import (
	context "context"

	models "github.com/cdumange/notion-htmx-go/models"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// mockTaskCreator is an autogenerated mock type for the taskCreator type
type mockTaskCreator struct {
	mock.Mock
}

// CreateTask provides a mock function with given fields: ctx, task
func (_m *mockTaskCreator) CreateTask(ctx context.Context, task models.Task) (uuid.UUID, error) {
	ret := _m.Called(ctx, task)

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Task) (uuid.UUID, error)); ok {
		return rf(ctx, task)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.Task) uuid.UUID); ok {
		r0 = rf(ctx, task)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.Task) error); ok {
		r1 = rf(ctx, task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newMockTaskCreator creates a new instance of mockTaskCreator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockTaskCreator(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockTaskCreator {
	mock := &mockTaskCreator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}