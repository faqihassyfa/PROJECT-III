package mysql

import (
	"PROJECT-III/config"
	productdata "PROJECT-III/features/Admin/data"
	orderdata "PROJECT-III/features/Order/data"
	userdata "PROJECT-III/features/User/data"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.Username,
		cfg.Password,
		cfg.Address,
		cfg.Port,
		cfg.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	return db
}

func MigrateData(db *gorm.DB) {
	db.AutoMigrate(userdata.User{})
	db.AutoMigrate(orderdata.Order{})
	db.AutoMigrate(productdata.Product{})
}
