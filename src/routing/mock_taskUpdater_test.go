// Code generated by mockery v2.43.2. DO NOT EDIT.

package routing

import (
	context "context"

	models "github.com/cdumange/notion-htmx-go/models"
	mock "github.com/stretchr/testify/mock"
)

// mockTaskUpdater is an autogenerated mock type for the taskUpdater type
type mockTaskUpdater struct {
	mock.Mock
}

// UpdateTask provides a mock function with given fields: ctx, task
func (_m *mockTaskUpdater) UpdateTask(ctx context.Context, task models.Task) error {
	ret := _m.Called(ctx, task)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Task) error); ok {
		r0 = rf(ctx, task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockTaskUpdater creates a new instance of mockTaskUpdater. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockTaskUpdater(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockTaskUpdater {
	mock := &mockTaskUpdater{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
