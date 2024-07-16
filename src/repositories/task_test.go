package repositories

import (
	"context"
	"testing"

	"github.com/cdumange/notion-htmx-go/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
)

func (s *postgresSuite) TestChangeCategory() {
	var err error
	t := s.T()
	ctx := context.Background()

	cleanTable(t, s.db)
	catR := NewCategoryRepository(s.db)
	taskR := NewTaskRepository(s.db)

	cat1 := models.Category{
		Title: "todo",
	}
	cat1.ID, err = catR.CreateCategory(ctx, cat1)
	require.NoError(t, err)

	cat2 := models.Category{
		Title: "done",
	}
	cat2.ID, err = catR.CreateCategory(ctx, cat2)
	require.NoError(t, err)

	task := models.Task{
		CategoryID: cat1.ID,
		Title:      "a task",
	}
	task.ID, err = taskR.CreateTask(ctx, task)
	require.NoError(t, err)

	require.NoError(t, taskR.ChangeCategory(ctx, task.ID, cat2.ID))

	v, err := taskR.GetTaskByCategory(ctx, cat1.ID)
	require.NoError(t, err)
	require.Empty(t, v)

	v, err = taskR.GetTaskByCategory(ctx, cat2.ID)
	require.NoError(t, err)
	require.Len(t, v, 1)
	require.Equal(t, v[0].ID, task.ID)
}

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
