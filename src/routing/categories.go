package routing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cdumange/notion-htmx-go/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type getCategoryWithTasksUC interface {
	GetCategoryWithTasks(ctx context.Context, categoryID uuid.UUID) (models.Category, error)
}

type getAllCategoryUC interface {
	GetCategoriesWithTasks(ctx context.Context) ([]models.Category, error)
}

func registerCategoryHandlers(app *echo.Echo, deps Dependencies) {
	group := app.Group("/categories")
	group.GET("/full/:categoryID", getCategoryWithTasks(deps.CategoryFullGetter))
	group.GET("/all", getAllCategories(deps.GetAllCategory))
}

func getCategoryWithTasks(uc getCategoryWithTasksUC) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.Param("categoryID"))
		categoryID, err := uuid.Parse(c.Param("categoryID"))
		if err != nil || categoryID == uuid.Nil {
			return c.NoContent(http.StatusBadRequest)
		}

		v, err := uc.GetCategoryWithTasks(c.Request().Context(), categoryID)
		switch err {
		case nil:
			return c.JSON(http.StatusOK, v)
		case models.ErrCategoryNotFound:
			return c.NoContent(http.StatusNotFound)
		default:
			c.Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}
}

func getAllCategories(s getAllCategoryUC) echo.HandlerFunc {
	return func(c echo.Context) error {
		v, err := s.GetCategoriesWithTasks(c.Request().Context())
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, v)
	}
}
