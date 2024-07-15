package routing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/cdumange/notion-htmx-go/models"
	"github.com/cdumange/notion-htmx-go/templates"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type categoryService interface {
	GetCategoryWithTasks(ctx context.Context, categoryID uuid.UUID) (models.Category, error)
	GetCategoriesWithTasks(ctx context.Context) ([]models.Category, error)
	CreateCategory(ctx context.Context, category models.Category) (uuid.UUID, error)
}

type taskService interface {
	UpdateTask(ctx context.Context, task models.Task) error
	CreateTask(ctx context.Context, task models.Task) (uuid.UUID, error)
	DeleteTask(ctx context.Context, ID uuid.UUID) error
	ChangeCategory(ctx context.Context, taskID, categoryID uuid.UUID) error
}

func registerCategoryHandlers(app *echo.Echo, deps Dependencies) {
	group := app.Group("/categories")
	group.GET("/full/:categoryID", getCategoryWithTasks(deps.CategoryService))
	group.GET("/all", getAllCategories(deps.CategoryService))
	group.POST("", createCategory(deps.CategoryService))
}

func createCategory(uc categoryService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var cat models.Category
		if err := c.Bind(&cat); err != nil ||
			len(cat.Title) == 0 {
			return c.NoContent(http.StatusBadRequest)
		}

		_, err := uc.CreateCategory(c.Request().Context(), cat)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		cats, err := uc.GetCategoriesWithTasks(c.Request().Context())
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		templ.Handler(templates.CategoryList(cats)).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
}

func getCategoryWithTasks(uc categoryService) echo.HandlerFunc {
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

func getAllCategories(s categoryService) echo.HandlerFunc {
	return func(c echo.Context) error {
		v, err := s.GetCategoriesWithTasks(c.Request().Context())
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, v)
	}
}
