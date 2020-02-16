package web

import (
	"github.com/labstack/echo"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", serveFile)
	e.GET("/:filename", serveFile)
}
