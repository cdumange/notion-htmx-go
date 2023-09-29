package routing

import (
	"net/http"

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

	CategoryFullGetter getCategoryWithTasksUC
	GetAllCategory     getAllCategoryUC
}

func index(deps Dependencies) echo.HandlerFunc {
	return func(c echo.Context) error {
		cats, err := deps.GetAllCategory.GetCategoriesWithTasks(c.Request().Context())
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.Render(http.StatusOK, "index.html", cats)
	}
}
