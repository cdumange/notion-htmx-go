// Code generated by mockery v2.33.3. DO NOT EDIT.

package routing

import (
	context "context"

	models "github.com/cdumange/notion-htmx-go/models"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// mockGetCategoryWithTasksUC is an autogenerated mock type for the getCategoryWithTasksUC type
type mockGetCategoryWithTasksUC struct {
	mock.Mock
}

// GetCategoryWithTasks provides a mock function with given fields: ctx, categoryID
func (_m *mockGetCategoryWithTasksUC) GetCategoryWithTasks(ctx context.Context, categoryID uuid.UUID) (models.Category, error) {
	ret := _m.Called(ctx, categoryID)

	var r0 models.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (models.Category, error)); ok {
		return rf(ctx, categoryID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) models.Category); ok {
		r0 = rf(ctx, categoryID)
	} else {
		r0 = ret.Get(0).(models.Category)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, categoryID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newMockGetCategoryWithTasksUC creates a new instance of mockGetCategoryWithTasksUC. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockGetCategoryWithTasksUC(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockGetCategoryWithTasksUC {
	mock := &mockGetCategoryWithTasksUC{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}