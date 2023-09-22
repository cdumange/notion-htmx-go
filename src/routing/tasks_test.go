package routing

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cdumange/notion-htmx-go/common"
	"github.com/cdumange/notion-htmx-go/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_createTask(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		api := common.NewEcho()
		rec := httptest.NewRecorder()

		task := models.Task{
			CategoryID: uuid.New(),
			Title:      "a title",
		}
		taskID := uuid.New()

		s := newMockTaskCreator(t)
		s.On("CreateTask", mock.Anything, task).Return(taskID, nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/tasks", toJSONReader(t, task))
		req.Header.Add("Content-Type", "application/json")
		c := api.NewContext(req, rec)

		assert.NoError(t, createTask(s)(c))
		assert.Equal(t, http.StatusCreated, rec.Result().StatusCode)

		ret := fromJSON[struct {
			ID uuid.UUID `json:"id"`
		}](t, rec.Result().Body)

		assert.Equal(t, taskID, ret.ID)
	})

	t.Run("bad request", func(t *testing.T) {
		api := common.NewEcho()
		rec := httptest.NewRecorder()
		task := models.Task{
			Title: "a title",
		}

		s := newMockTaskCreator(t)

		req := httptest.NewRequest(http.MethodPost, "/tasks", toJSONReader(t, task))
		req.Header.Add("Content-Type", "application/json")
		c := api.NewContext(req, rec)

		assert.NoError(t, createTask(s)(c))
		assert.Equal(t, http.StatusBadRequest, rec.Result().StatusCode)
	})

	t.Run("500", func(t *testing.T) {
		api := common.NewEcho()
		rec := httptest.NewRecorder()

		task := models.Task{
			CategoryID: uuid.New(),
			Title:      "a title",
		}

		s := newMockTaskCreator(t)
		s.On("CreateTask", mock.Anything, task).Return(uuid.Nil, errors.New("an error")).Once()

		req := httptest.NewRequest(http.MethodPost, "/tasks", toJSONReader(t, task))
		req.Header.Add("Content-Type", "application/json")
		c := api.NewContext(req, rec)

		assert.NoError(t, createTask(s)(c))
		assert.Equal(t, http.StatusInternalServerError, rec.Result().StatusCode)
	})
}

func toJSONReader[T any](t *testing.T, body T) io.Reader {
	t.Helper()

	s, err := json.Marshal(body)
	require.NoError(t, err)

	return bytes.NewReader(s)
}

func fromJSON[T any](t *testing.T, buffer io.Reader) T {
	var v T

	b, err := io.ReadAll(buffer)
	require.NoError(t, err)

	require.NoError(t, json.Unmarshal(b, &v))

	return v
}
