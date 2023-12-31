package repositories

import (
	"context"
	"testing"

	"github.com/cdumange/notion-htmx-go/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func TestCreateTask(t *testing.T) {
	ctx := context.Background()
	categoryID := uuid.New()
	title := "a title"

	db := initDB(t)

	s := NewTaskRepository(db)

	ID, err := s.CreateTask(ctx, models.Task{
		CategoryID: categoryID,
		Title:      title,
	})

	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, ID)

	r := db.QueryRowContext(ctx, "SELECT title FROM tasks WHERE id = $1", ID)
	assert.NoError(t, r.Err())

	var retTitle string
	assert.NoError(t, r.Scan(&retTitle))
	assert.Equal(t, title, retTitle)
}

func TestGetTaskByCategory(t *testing.T) {
	ctx := context.Background()
	db := initDB(t)
	s := NewTaskRepository(db)

	t.Run("not exists", func(t *testing.T) {
		cleanTable(t, db)
		v, err := s.GetTaskByCategory(ctx, uuid.New())
		assert.NoError(t, err)
		assert.Empty(t, v)
	})

	t.Run("exists", func(t *testing.T) {
		categoryID := uuid.New()
		taskID, err := s.CreateTask(ctx, models.Task{
			CategoryID: categoryID,
			Title:      "a title",
		})

		assert.NoError(t, err)

		v, err := s.GetTaskByCategory(ctx, categoryID)
		assert.NoError(t, err)
		assert.NotEmpty(t, v)

		assert.Equal(t, taskID, v[0].ID)
	})
}
