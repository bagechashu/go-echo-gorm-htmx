package templates

import (
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.StaticFS("/", assetsFS)
}
