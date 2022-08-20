package delivery

import (
	"PROJECT-III/config"
	"PROJECT-III/domain"
	"PROJECT-III/features/Admin/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteProducts(e *echo.Echo, ah domain.AdminHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.Pre(middleware.RemoveTrailingSlash())
	e.Pre(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	product := e.Group("/admins")
	product.PUT("/:productid", ah.Update(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	product.DELETE("/:productid", ah.Delete(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	product.POST("", ah.Create(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	product.GET("", ah.ReadAll(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	// product.GET("/history", ah.ReadHistory(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
