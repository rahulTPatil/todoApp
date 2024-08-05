package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, `<h2>This is best</h2>`)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
