package routing

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/cdumange/notion-htmx-go/models"
	"github.com/cdumange/notion-htmx-go/templates/mixins"
)

func registerTasksEndpoint(app *echo.Echo, deps Dependencies) {
	group := app.Group("tasks")
	group.PUT("/id/:id", updateTask(deps.TaskRepository))
	group.PUT("/category/:id", updateTask(deps.TaskRepository))
	group.PUT("/cat", updateTaskCat(deps.TaskRepository))

	group.POST("", createTask(deps.TaskRepository))
	group.DELETE("/:id", deleteTask(deps.TaskRepository))
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

		templ.Handler(mixins.Task(task)).ServeHTTP(c.Response().Writer, c.Request())
		return nil
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
		templ.Handler(mixins.Task(task)).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
}

type taskDeletor interface {
	DeleteTask(ctx context.Context, ID uuid.UUID) error
}

func deleteTask(deleter taskDeletor) echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		err = deleter.DeleteTask(c.Request().Context(), ID)
		if err != nil {
			c.Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.NoContent(http.StatusAccepted)
	}
}

type taskCategoriser interface {
	ChangeCategory(ctx context.Context, taskID, categoryID uuid.UUID) error
}

type updateTaskInput struct {
	CategoryID uuid.UUID `form:"category_id" validate:"required"`
	TaskID     uuid.UUID `form:"task_id" validate:"required"`
}

func updateTaskCat(cat taskCategoriser) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input updateTaskInput
		if err := c.Bind(&input); err != nil {
			return c.NoContent(http.StatusNoContent)
		}

		if err := c.Validate(input); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		if err := cat.ChangeCategory(c.Request().Context(), input.TaskID, input.CategoryID); err != nil {
			c.Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusAccepted)
	}
}
