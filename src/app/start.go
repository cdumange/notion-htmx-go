package app

import (
	"database/sql"

	"github.com/cdumange/notion-htmx-go/common"
	"github.com/cdumange/notion-htmx-go/repositories"
	"github.com/cdumange/notion-htmx-go/routing"
	"github.com/cdumange/notion-htmx-go/templates"
)

func Start() {
	e := common.NewEcho()

	db, err := sql.Open("postgres", "postgresql://postgres:postgres@localhost/public?sslmode=disable")
	if err != nil {
		panic(err)
	}

	taskRepo := repositories.NewTaskRepository(db)

	templates.RegisterTemplates(e)
	routing.LoadRouter(e, routing.Dependencies{
		TaskCreator: taskRepo,
	})
	e.Logger.Panic(e.Start(":3000"))
}