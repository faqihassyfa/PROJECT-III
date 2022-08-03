package factory

import (
	"PROJECT-III/config"
	awss3 "PROJECT-III/infrastructure/database/aws"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	ud "PROJECT-III/features/User/data"
	udeli "PROJECT-III/features/User/delivery"
	uc "PROJECT-III/features/User/usecase"

	ad "PROJECT-III/features/Admin/data"
	adeli "PROJECT-III/features/Admin/delivery"
	ac "PROJECT-III/features/Admin/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB, cfg config.AppConfig) {
	validator := validator.New()

	userData := ud.New(db)
	useCase := uc.New(userData, validator)
	userHandler := udeli.New(useCase, userData)
	udeli.RouteUser(e, userHandler)

	adminData := ad.New(db)
	adminCase := ac.New(adminData, validator)
	adminHandler := adeli.New(adminData, adminCase, awss3.InitS3(cfg.Keys3, cfg.Secrets3, cfg.Regions3))
	adeli.RouteProducts(e, adminHandler)
}
