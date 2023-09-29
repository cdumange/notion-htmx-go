package routing

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/cdumange/notion-htmx-go/models"
)

func registerTasksEndpoint(app *echo.Echo, deps Dependencies) {
	group := app.Group("tasks")
	group.PUT("/id/:id", updateTask(deps.TaskUpdater))
	group.PUT("/category/:id", updateTask(deps.TaskUpdater))

	group.POST("", createTask(deps.TaskCreator))
}

type taskUpdater interface {
	UpdateTask(ctx context.Context, task models.Task) error
}

func updateTask(s taskUpdater) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var task models.Task
		if err := c.Bind(&task); err != nil && task.CategoryID != uuid.Nil && len(task.Title) > 0 {
			return c.NoContent(http.StatusBadRequest)
		}

		if err = c.Validate(task); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		if task.ID, err = uuid.Parse(c.Param("id")); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		if err = s.UpdateTask(c.Request().Context(), task); err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.Render(http.StatusAccepted, "task.html", task)
	}
}

type taskCreator interface {
	CreateTask(ctx context.Context, task models.Task) (uuid.UUID, error)
}

func createTask(creator taskCreator) echo.HandlerFunc {
	return func(c echo.Context) error {
		var task models.Task
		if err := c.Bind(&task); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		if err := c.Validate(task); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		ID, err := creator.CreateTask(c.Request().Context(), task)
		if err != nil {
			c.Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}

		task.ID = ID
		return c.Render(http.StatusCreated, "task.html", task)
	}
}
