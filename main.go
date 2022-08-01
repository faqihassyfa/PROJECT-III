package main

import (
	"PROJECT-III/config"
	"PROJECT-III/factory"
	"PROJECT-III/infrastructure/database/mysql"
	"PROJECT-III/migration"
	"fmt"

	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	migration.InitMigrate(db)
	e := echo.New()
	// e.Use(middleware.CORS())

	factory.InitFactory(e, db)

	fmt.Println("application is running ....")
	dsn := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(dsn))
}
