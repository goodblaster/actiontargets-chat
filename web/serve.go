package web

import (
	"path/filepath"

	"github.com/labstack/echo"
)

const webPath = "web/files"

func serveFile(c echo.Context) error {
	return c.File(filepath.Join(webPath, c.Param("filename")))
}
