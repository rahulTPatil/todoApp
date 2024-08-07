package main

import (
	"html/template"
	"io"

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

type Todo struct {
	Name   string
	Added  string
	Status bool
}

func newTodo(Name string, Added string, Status bool) Todo {
	return Todo{
		Name:   Name,
		Added:  Added,
		Status: Status,
	}
}

type Todos = []Todo

type Data struct {
	Todos Todos
}

func newData() Data {
	return Data{
		Todos: []Todo{
			newTodo("Try GPG", "2024-08-07", false),
			newTodo("Try ChatGPT", "2024-10-20", false),
		},
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	data := newData()

	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", data)
	})

	e.POST("/addTodo", func(c echo.Context) error {
		d := c.FormValue("todoName")
		data.Todos = append(data.Todos, newTodo(d, "2024-08-07", false))
		return c.Render(200, "table", data)
	})

	e.Logger.Fatal(e.Start(":1323"))

}
