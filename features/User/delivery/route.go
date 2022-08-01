package delivery

import (
	"PROJECT-III/domain"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteUser(e *echo.Echo, uh domain.UserHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))
}
