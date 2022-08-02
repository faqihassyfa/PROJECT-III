package migration

import (
	_mOrder "PROJECT-III/features/Order/data"
	_mUsers "PROJECT-III/features/User/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_mUsers.User{})
	db.AutoMigrate(&_mOrder.Order{})
}
