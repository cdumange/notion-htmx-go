package routing

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cdumange/notion-htmx-go/common"
	"github.com/cdumange/notion-htmx-go/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_getCategoryWithTasks(t *testing.T) {
	api := common.NewEcho()

	categoryID := uuid.New()

	t.Run("200", func(t *testing.T) {
		expectedCategory := models.Category{
			ID:    categoryID,
			Title: "a title",
			List: []models.Task{
				{
					ID:         uuid.New(),
					CategoryID: categoryID,
					Title:      "a title again",
				},
			},
		}

		rec := httptest.NewRecorder()

		mockS := newMockGetCategoryWithTasksUC(t)
		mockS.On("GetCategoryWithTasks", mock.Anything, categoryID).Return(expectedCategory, nil).Once()

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/full/%s", categoryID.String()), nil)
		c := api.NewContext(req, rec)
		c.SetParamNames("categoryID")
		c.SetParamValues(categoryID.String())

		assert.NoError(t, getCategoryWithTasks(mockS)(c))
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

		v := fromJSON[models.Category](t, rec.Result().Body)
		assert.Equal(t, expectedCategory, v)
	})

	t.Run("400", func(t *testing.T) {
		rec := httptest.NewRecorder()

		mockS := newMockGetCategoryWithTasksUC(t)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/full/%s", categoryID.String()), nil)
		c := api.NewContext(req, rec)
		c.SetParamNames("categoryID")

		assert.NoError(t, getCategoryWithTasks(mockS)(c))
		assert.Equal(t, http.StatusBadRequest, rec.Result().StatusCode)
	})

	t.Run("500", func(t *testing.T) {
		expectedCategory := models.Category{
			ID:    categoryID,
			Title: "a title",
			List: []models.Task{
				{
					ID:         uuid.New(),
					CategoryID: categoryID,
					Title:      "a title again",
				},
			},
		}

		rec := httptest.NewRecorder()

		mockS := newMockGetCategoryWithTasksUC(t)
		mockS.On("GetCategoryWithTasks", mock.Anything, categoryID).
			Return(expectedCategory, errors.New("an error")).
			Once()

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/full/%s", categoryID.String()), nil)
		c := api.NewContext(req, rec)
		c.SetParamNames("categoryID")
		c.SetParamValues(categoryID.String())

		assert.NoError(t, getCategoryWithTasks(mockS)(c))
		assert.Equal(t, http.StatusInternalServerError, rec.Result().StatusCode)
	})

	t.Run("404", func(t *testing.T) {
		expectedCategory := models.Category{
			ID:    categoryID,
			Title: "a title",
			List: []models.Task{
				{
					ID:         uuid.New(),
					CategoryID: categoryID,
					Title:      "a title again",
				},
			},
		}

		rec := httptest.NewRecorder()

		mockS := newMockGetCategoryWithTasksUC(t)
		mockS.On("GetCategoryWithTasks", mock.Anything, categoryID).
			Return(expectedCategory, models.ErrCategoryNotFound).
			Once()

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/full/%s", categoryID.String()), nil)
		c := api.NewContext(req, rec)
		c.SetParamNames("categoryID")
		c.SetParamValues(categoryID.String())

		assert.NoError(t, getCategoryWithTasks(mockS)(c))
		assert.Equal(t, http.StatusNotFound, rec.Result().StatusCode)
	})
}
