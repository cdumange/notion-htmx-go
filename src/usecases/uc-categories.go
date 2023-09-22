package usecases

import (
	"context"

	"github.com/cdumange/notion-htmx-go/models"
	"github.com/google/uuid"
)

type categoryRepository interface {
	GetCategory(ctx context.Context, ID uuid.UUID) (models.Category, error)
	CreateCategory(ctx context.Context, category models.Category) (uuid.UUID, error)
}

type taskRepository interface {
	GetTaskByCategory(ctx context.Context, categoryID uuid.UUID) ([]models.Task, error)
}

type CategoriesUsecases struct {
	categories categoryRepository
	tasks      taskRepository
}

func NewCategoriesUsecases(
	categories categoryRepository,
	tasks taskRepository,
) *CategoriesUsecases {
	return &CategoriesUsecases{
		categories: categories,
		tasks:      tasks,
	}
}

// GetCategory retrieves a category. Task will remain empty.
func (uc *CategoriesUsecases) GetCategory(ctx context.Context, categoryID uuid.UUID) (models.Category, error) {
	return uc.categories.GetCategory(ctx, categoryID)
}

// GetCategoryWithTasks retrieves a category with its tasks.
func (uc *CategoriesUsecases) GetCategoryWithTasks(ctx context.Context, categoryID uuid.UUID) (models.Category, error) {
	cat, err := uc.categories.GetCategory(ctx, categoryID)
	if err != nil {
		return models.Category{}, err
	}

	cat.List, err = uc.tasks.GetTaskByCategory(ctx, categoryID)
	if err != nil {
		return models.Category{}, err
	}

	return cat, nil
}
