package data

import (
	"PROJECT-III/domain"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderID    string
	UserID     int
	ProductID  int
	Price      int
	Qty        int
	Totalprice int
	Status     int
	Payment    string
}

func (o *Order) ToModel() domain.Order {
	return domain.Order{
		ID:         int(o.ID),
		OrderID:    o.OrderID,
		UserID:     o.UserID,
		ProductID:  o.ProductID,
		Price:      o.Price,
		Qty:        o.Qty,
		Totalprice: o.Totalprice,
		Status:     o.Status,
		Payment:    o.Payment,
		CreatedAt:  o.CreatedAt,
	}
}

func FromModel(data domain.Order) Order {
	var res Order
	res.ID = uint(data.ID)
	res.OrderID = data.OrderID
	res.UserID = data.UserID
	res.ProductID = data.ProductID
	res.Price = data.Price
	res.Qty = data.Qty
	res.Totalprice = data.Totalprice
	res.Status = data.Status
	res.Payment = data.Payment
	res.CreatedAt = data.CreatedAt

	return res
}

func ParseToArr(arr []Order) []domain.Order {
	var res []domain.Order

	for _, val := range arr {
		res = append(res, val.ToModel())
	}

	return res
}
