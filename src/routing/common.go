package routing

import (
	"net/http"

	"github.com/cdumange/notion-htmx-go/models"

	"github.com/labstack/echo/v4"
)

func LoadRouter(app *echo.Echo, deps Dependencies) {
	app.GET("", index)

	registerTasksEndpoint(app, deps)
}

type Dependencies struct {
	TaskCreator taskCreator

	CategoryFullGetter getCategoryWithTasksUC
}

func index(ctx echo.Context) error {
	var err error

	if err = ctx.Render(http.StatusOK, "index.html", []models.Category{}); err != nil {
		ctx.Logger().Error(err)
	}

	return err
}
