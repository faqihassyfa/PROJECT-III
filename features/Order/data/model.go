package data

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID            int
	Userid        int
	PaymentMethod string
	TotalQty      int
	Totalprice    int
	Status        int
}
