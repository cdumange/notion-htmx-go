package app

import (
	"github.com/cdumange/notion-htmx-go/routing"
	"github.com/cdumange/notion-htmx-go/templates"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Start() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if err = next(c); err != nil {
				c.Logger().Error(err)
			}
			return err
		}
	})	


	templates.RegisterTemplates(e)
	routing.LoadRouter(e)
	e.Logger.Panic(e.Start(":3000"))
}