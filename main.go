package main

import (
	"PROJECT-III/config"
	"PROJECT-III/infrastructure/database/mysql"
	"PROJECT-III/migration"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	migration.InitMigrate(db)
	e := echo.New()
	e.Use(middleware.CORS())

	fmt.Println("application is running ....")
	dsn := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(dsn))
}
