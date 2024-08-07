package main

import (
	"fmt"
	"html/template"
	"io"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("*.html")),
	}
}

type Table struct {
	Name   string
	Added  string
	Status bool
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	data := Table{Name: "Try Go!", Added: "2024-08-24", Status: true}

	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", data)
	})

	e.POST("/addTodo", func(c echo.Context) error {
		d := c.Param("todoName")
		data2 := Table{Name: d, Added: time.Now().String(), Status: false}
		fmt.Println(data2)
		return c.Render(200, "table", data2)
	})

	e.Logger.Fatal(e.Start(":1323"))

}
