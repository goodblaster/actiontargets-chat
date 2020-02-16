package web

import (
	"strconv"

	"github.com/labstack/echo"
)

// Run - Panic instead of returning error because this
// should be run in a go routine.
func Run(port int) {
	e := echo.New()
	InitRoutes(e)
	err := e.Start(":" + strconv.Itoa(port))
	if err != nil {
		panic("Start Web Server: " + err.Error())
	}
}
