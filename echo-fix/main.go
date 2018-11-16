package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RewriteWithConfig(middleware.RewriteConfig{
		Rules: map[string]string{
			"/api/*/mgmt/proj/*/agt": "/api/$1/hosts/$2/",
			"/api/*/mgmt/proj":       "/api/$1/eng/",
		},
	}))
	e.GET("/api/:version/hosts/:name/", func(c echo.Context) error {
		return c.String(200, "hosts")
	})
	e.GET("/api/:verson/eng/", func(c echo.Context) error {
		return c.String(200, "eng")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
