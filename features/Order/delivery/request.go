package delivery

import "PROJECT-III/domain"

type OrderFormat struct {
	Productid int `json:"productid" form:"productid" validate:"required"`
	Qty       int `json:"qty" form:"qty" validate:"required"`
}

func (of *OrderFormat) ToModel() domain.Order {
	return domain.Order{
		Productid: of.Productid,
		Qty:       of.Qty,
	}
}
