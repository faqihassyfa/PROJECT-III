package data

import (
	"PROJECT-III/domain"
	"log"

	"gorm.io/gorm"
)

type orderData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.OrderData {
	return &orderData{
		db: db,
	}
}

func (od *orderData) CreateOrderData(neworder domain.Order) domain.Order {
	var order = FromModel(neworder)

	err := od.db.Create(&order)

	if err.Error != nil {
		log.Println("cant create order", err.Error)
		return domain.Order{}
	}
	return order.ToModel()

}
