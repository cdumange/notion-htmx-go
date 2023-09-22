package routing

import (
	"net/http"

	"github.com/cdumange/notion-htmx-go/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func LoadRouter(app *echo.Echo) {
	app.GET("",index )

	registerTasksEndpoint(app)
}

func index(ctx echo.Context) error {
	var err error
	
	if err = ctx.Render(http.StatusOK, "index.html", []models.Categories {
		{
			ID: uuid.New(),
			Title: "TODO",
			List: models.TaskList{
				Tasks: []models.Task {
					{
						ID: uuid.New(),
						Title: "retest",
					},
					{
						ID: uuid.New(),
						Title: "test",
					},
				},
			},
		},
		{
			ID: uuid.New(),
			Title: "Doing",
			List: models.TaskList{
				Tasks: []models.Task {
					{
						ID: uuid.New(),
						Title: "dev",
					},
					{
						ID: uuid.New(),
						Title: "TU",
					},
				},
			},
		},
	}); err != nil {
		ctx.Logger().Error(err)
	}

	return err
}