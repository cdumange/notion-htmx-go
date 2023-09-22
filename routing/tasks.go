package routing

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/cdumange/notion-htmx-go/models"
)

func registerTasksEndpoint(app *echo.Echo) {
	group := app.Group("tasks")
	group.PUT("/id/:id", updateTask)
	group.PUT("/category/:id", updateTask)
}

func updateTask(c echo.Context) (err error) {
	var task models.Task
	if err := c.Bind(&task); err != nil && task.CategoryID != uuid.Nil && len(task.Title) > 0  {
		return c.NoContent(http.StatusBadRequest)
	}
	
	if task.ID, err = uuid.Parse(c.Param("id")); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	c.Logger().Debugf("received : %v", task)

	return c.Render(http.StatusAccepted, "task.html", task)
}