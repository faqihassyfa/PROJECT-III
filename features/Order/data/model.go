package data

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Userid     int
	TotalQty   int
	Totalprice int
	Status     int
}
