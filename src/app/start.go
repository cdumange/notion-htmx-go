package app

import (
	"database/sql"

	"github.com/cdumange/notion-htmx-go/common"
	"github.com/cdumange/notion-htmx-go/repositories"
	"github.com/cdumange/notion-htmx-go/routing"
	"github.com/cdumange/notion-htmx-go/templates"
	"github.com/cdumange/notion-htmx-go/usecases"

	_ "github.com/lib/pq"
)

func Start() {
	e := common.NewEcho()

	db, err := sql.Open("postgres", "postgresql://postgres:postgres@localhost/public?sslmode=disable")
	if err != nil {
		panic(err)
	}

	taskRepo := repositories.NewTaskRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)

	ucRepositories := usecases.NewCategoriesUsecases(categoryRepo, taskRepo)

	templates.RegisterTemplates(e)
	routing.LoadRouter(e, routing.Dependencies{
		TaskCreator:        taskRepo,
		TaskUpdater:        taskRepo,
		TaskDeletor:        taskRepo,
		GetAllCategory:     ucRepositories,
		CategoryFullGetter: ucRepositories,
	})
	e.Logger.Panic(e.Start(":3000"))
}
