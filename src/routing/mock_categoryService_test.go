// Code generated by mockery v2.43.2. DO NOT EDIT.

package routing

import (
	context "context"

	models "github.com/cdumange/notion-htmx-go/models"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// mockCategoryService is an autogenerated mock type for the categoryService type
type mockCategoryService struct {
	mock.Mock
}

// CreateCategory provides a mock function with given fields: ctx, category
func (_m *mockCategoryService) CreateCategory(ctx context.Context, category models.Category) (uuid.UUID, error) {
	ret := _m.Called(ctx, category)

	if len(ret) == 0 {
		panic("no return value specified for CreateCategory")
	}

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Category) (uuid.UUID, error)); ok {
		return rf(ctx, category)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.Category) uuid.UUID); ok {
		r0 = rf(ctx, category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.Category) error); ok {
		r1 = rf(ctx, category)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategoriesWithTasks provides a mock function with given fields: ctx
func (_m *mockCategoryService) GetCategoriesWithTasks(ctx context.Context) ([]models.Category, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetCategoriesWithTasks")
	}

	var r0 []models.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]models.Category, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []models.Category); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategoryWithTasks provides a mock function with given fields: ctx, categoryID
func (_m *mockCategoryService) GetCategoryWithTasks(ctx context.Context, categoryID uuid.UUID) (models.Category, error) {
	ret := _m.Called(ctx, categoryID)

	if len(ret) == 0 {
		panic("no return value specified for GetCategoryWithTasks")
	}

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

// newMockCategoryService creates a new instance of mockCategoryService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockCategoryService(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockCategoryService {
	mock := &mockCategoryService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}