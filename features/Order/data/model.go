package data

import (
	"PROJECT-III/domain"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Userid        int
	PaymentMethod string
	Qty           int
	Totalprice    int
	Status        int
	Productid     int
	CreatedAt     time.Time
}

func (o *Order) ToModel() domain.Order {
	return domain.Order{
		ID:            int(o.ID),
		Userid:        o.Userid,
		PaymentMethod: o.PaymentMethod,
		Qty:           o.Qty,
		Totalprice:    o.Totalprice,
		Status:        o.Status,
		Productid:     o.Productid,
		CreatedAt:     o.CreatedAt,
	}
}

func ParseToArr(arr []Order) []domain.Order {
	var res []domain.Order

	for _, val := range arr {
		res = append(res, val.ToModel())
	}

	return res
}

func FromModel(data domain.Order) Order {
	var res Order
	res.ID = uint(data.ID)
	res.Userid = data.Userid
	res.PaymentMethod = data.PaymentMethod
	res.Qty = data.Qty
	res.Totalprice = data.Totalprice
	res.Status = data.Status
	res.Productid = data.Productid
	res.CreatedAt = data.CreatedAt

	return res
}
