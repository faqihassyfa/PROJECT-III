package delivery

import (
	"PROJECT-III/config"
	"PROJECT-III/domain"
	"PROJECT-III/features/User/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteUser(e *echo.Echo, uh domain.UserHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET},
	}))
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	login := e.Group("/login")
	user := e.Group("/users")
	// user.GET("", uh.Account(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	user.PUT("", uh.Update(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	user.POST("", uh.Register())
	user.DELETE("", uh.Delete(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	login.POST("", uh.Login())

	e.GET("/products", uh.Product())
}
