package routing

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/cdumange/notion-htmx-go/templates"
	"github.com/labstack/echo/v4"
)

// LoadRouter loads all the endpoints for the API.
func LoadRouter(app *echo.Echo, deps Dependencies) {
	app.GET("", index(deps))

	registerTasksEndpoint(app, deps)
}

// Dependencies holds the routing dependencies.
type Dependencies struct {
	TaskCreator taskCreator
	TaskUpdater taskUpdater
	TaskDeletor taskDeletor

	CategoryFullGetter getCategoryWithTasksUC
	GetAllCategory     getAllCategoryUC
}

func index(deps Dependencies) echo.HandlerFunc {
	return func(c echo.Context) error {
		cats, err := deps.GetAllCategory.GetCategoriesWithTasks(c.Request().Context())
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		templ.Handler(templates.Index(cats)).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
}
