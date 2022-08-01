package factory

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	ud "PROJECT-III/features/User/data"
	udeli "PROJECT-III/features/User/delivery"
	uc "PROJECT-III/features/User/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	validator := validator.New()

	userData := ud.New(db)
	useCase := uc.New(userData, validator)
	userHandler := udeli.New(useCase, userData)
	udeli.RouteUser(e, userHandler)
}
