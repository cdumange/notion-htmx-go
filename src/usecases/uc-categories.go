package usecases

import (
	"context"
	"fmt"
	"sync"

	"github.com/cdumange/notion-htmx-go/models"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
)

type categoryRepository interface {
	GetCategory(ctx context.Context, ID uuid.UUID) (models.Category, error)
	CreateCategory(ctx context.Context, category models.Category) (uuid.UUID, error)
	GetAll(ctx context.Context) ([]models.Category, error)
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

// GetCategoriesWithTasks retrieves a category with its tasks.
func (uc *CategoriesUsecases) GetCategoriesWithTasks(ctx context.Context) ([]models.Category, error) {
	cats, err := uc.categories.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	errorChannel := make(chan error, 1)
	wg := new(sync.WaitGroup)

	for i, v := range cats {
		cat := v
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			var err error
			cat.List, err = uc.tasks.GetTaskByCategory(ctx, cat.ID)
			if err != nil {
				errorChannel <- fmt.Errorf("error while fetching tasks for %s: %w", cat.ID, err)
			}
			cats[i] = cat
		}(i)
	}
	wg.Wait()
	close(errorChannel)

	err = nil
	for e := range errorChannel {
		err = multierror.Append(err, e)
	}

	if err != nil {
		return nil, err
	}

	return cats, nil
}
