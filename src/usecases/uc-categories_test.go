package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/cdumange/notion-htmx-go/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_GetCategoryWithTasks(t *testing.T) {
	ctx := context.Background()
	ID := uuid.New()

	t.Run("ok", func(t *testing.T) {
		expectedCat := models.Category{
			ID:    ID,
			Title: "a title",
		}

		tasks := []models.Task{
			{
				ID:         uuid.New(),
				CategoryID: ID,
				Title:      "a task",
			},
			{
				ID:         uuid.New(),
				CategoryID: ID,
				Title:      "another task",
			},
		}
		mockCat := newMockCategoryRepository(t)
		mockCat.On("GetCategory", mock.Anything, ID).Return(expectedCat, nil).Once()

		mockTask := newMockTaskRepository(t)
		mockTask.On("GetTaskByCategory", mock.Anything, ID).Return(tasks, nil).Once()

		uc := CategoriesUsecases{
			mockCat,
			mockTask,
		}

		cat, err := uc.GetCategoryWithTasks(ctx, ID)
		assert.NoError(t, err)

		assert.Equal(t, ID, cat.ID)
		assert.Equal(t, expectedCat.Title, cat.Title)
		assert.Equal(t, tasks, cat.List)
	})

	t.Run("err on cat", func(t *testing.T) {
		mockCat := newMockCategoryRepository(t)
		mockCat.On("GetCategory", mock.Anything, ID).Return(models.Category{}, errors.New("an error")).Once()

		mockTask := newMockTaskRepository(t)

		uc := CategoriesUsecases{
			mockCat,
			mockTask,
		}

		_, err := uc.GetCategoryWithTasks(ctx, ID)
		assert.Error(t, err)
	})

	t.Run("err on tasks", func(t *testing.T) {
		expectedCat := models.Category{
			ID:    ID,
			Title: "a title",
		}

		mockCat := newMockCategoryRepository(t)
		mockCat.On("GetCategory", mock.Anything, ID).Return(expectedCat, nil).Once()

		mockTask := newMockTaskRepository(t)
		mockTask.On("GetTaskByCategory", mock.Anything, ID).Return(nil, errors.New("an error")).Once()

		uc := CategoriesUsecases{
			mockCat,
			mockTask,
		}

		_, err := uc.GetCategoryWithTasks(ctx, ID)
		assert.Error(t, err)
	})
}
