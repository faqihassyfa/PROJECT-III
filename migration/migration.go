package migration

import (
	_mUsers "PROJECT-III/domain"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_mUsers.User{})
}
