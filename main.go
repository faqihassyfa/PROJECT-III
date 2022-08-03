package main

import (
	"PROJECT-III/config"
	"PROJECT-III/factory"
	"PROJECT-III/infrastructure/database/mysql"
	"fmt"

	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateData(db)
	e := echo.New()
	// e.Use(middleware.CORS())

	factory.InitFactory(e, db, *cfg)

	fmt.Println("application is running ....")
	dsn := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(dsn))
}
