package templates

import (
	"html/template"
	"io"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

const templateExt = "html"

type templateRegister struct {
	templates *template.Template
}

func (t *templateRegister) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func RegisterTemplates(app *echo.Echo) {
	t := template.New("")
	if err := filepath.Walk("./templates", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() || !strings.HasSuffix(path, templateExt) {
			return nil
		}

		if _, err := t.ParseFiles(path); err != nil {
			return err
		}

		return nil
	}); err != nil {
		panic(err)
	}

	app.Renderer = &templateRegister{
		templates: t,
	}

	app.Static("/assets/", "assets")
}
