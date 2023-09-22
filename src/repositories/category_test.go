package repositories

import (
	"context"
	"testing"

	"github.com/cdumange/notion-htmx-go/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CreateCategory(t *testing.T) {
	db := initDB(t)

	cat := models.Category{
		Title: "a title",
	}

	var s = NewCategoryRepository(db)
	ID, err := s.CreateCategory(context.Background(), cat)
	assert.NoError(t, err)

	var testTitle string
	r := db.QueryRow("SELECT title FROM categories WHERE id = $1", ID)
	assert.NoError(t, r.Err())
	assert.NoError(t, r.Scan(&testTitle))

	assert.Equal(t, cat.Title, testTitle)
}

func Test_GetCategory(t *testing.T) {
	ctx := context.Background()
	cat := models.Category{
		Title: "another title",
	}
	db := initDB(t)
	s := NewCategoryRepository(db)

	t.Run("exists", func(t *testing.T) {
		cleanTable(t, db)

		ID, err := s.CreateCategory(ctx, cat)
		require.NoError(t, err)

		v, err := s.GetCategory(ctx, ID)
		assert.NoError(t, err)
		assert.Equal(t, ID, v.ID)
		assert.Equal(t, cat.Title, v.Title)
	})

	t.Run("not exists", func(t *testing.T) {
		cleanTable(t, db)

		_, err := s.GetCategory(ctx, uuid.New())
		assert.ErrorIs(t, err, models.ErrCategoryNotFound)
	})
}
