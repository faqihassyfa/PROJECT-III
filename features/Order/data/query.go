package data

import (
	"PROJECT-III/domain"
	"PROJECT-III/features/User/data"
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

func (od *orderData) TotalPriceOrder(neworder domain.Order) (orderprice, ordertotalprice int) {
	var proder domain.Product
	res2 := od.db.Where("ID = ?", neworder.ProductID).First(&proder)
	if res2.Error != nil {
		log.Println("Cannot retrieve object", res2.Error.Error())
		return 0, 0
	}

	sum := proder.Price * neworder.Qty

	neworder.Price = proder.Price

	return proder.Price, sum
}

func (od *orderData) CekStock(neworder domain.Order) bool {
	var proder domain.Product
	res2 := od.db.Where("ID = ?", neworder.ProductID).First(&proder)

	if res2.Error != nil {
		log.Println("Cannot retrieve object", res2.Error.Error())
		return false
	}

	cekstock := proder.Stock - neworder.Qty

	if cekstock > 0 {
		res3 := od.db.Model(&data.Product{}).Where("ID = ?", neworder.ProductID).Updates(data.Product{Stock: cekstock})
		if res3.Error != nil {
			log.Println("Cannot retrieve object", res3.Error.Error())
			return false
		}
	} else {
		log.Println("not enough stock")
		return false
	}

	return true
}

func (od *orderData) InsertOrderData(neworder domain.Order) domain.Order {
	order := FromModel(neworder)
	result := od.db.Create(&order)

	if result.Error != nil {
		log.Println("cant create order", result.Error.Error())
		return domain.Order{}
	}
	return order.ToModel()
}

func (od *orderData) ConfirmStatusData(orderid int) (row int) {
	result := od.db.Model(&Order{}).Where("id = ?", orderid).Update("status", "Confirmed")
	if result.Error != nil {
		log.Println("cant update data", result.Error.Error())
		return 0
	}
	if result.RowsAffected != 1 {
		return 0
	}
	return int(result.RowsAffected)
}

func (od *orderData) CancelStatusData(orderid int) (row int) {
	result := od.db.Model(&Order{}).Where("id = ?", orderid).Update("status", 9)
	if result.Error != nil {
		return 0
	}
	if result.RowsAffected != 1 {
		return 0
	}
	return int(result.RowsAffected)
}

func (od *orderData) ReturnStock(cancel domain.Order) {
	var proder domain.Product
	res2 := od.db.Where("ID = ?", cancel.ProductID).First(&proder)

	if res2.Error != nil {
		log.Println("Cannot retrieve object", res2.Error.Error())
	}

	var ordercanc Order
	res := od.db.Where("ID = ?", cancel.ProductID).First(&ordercanc)

	if res.Error != nil {
		log.Println("Cannot retrieve object", res.Error.Error())
	}

	stock := proder.Stock + cancel.Qty

	res3 := od.db.Model(&data.Product{}).Where("ID = ?", cancel.ProductID).Updates(data.Product{Stock: stock})
	if res3.Error != nil {
		log.Println("Cannot retrieve object", res3.Error.Error())
	}
}
