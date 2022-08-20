package delivery

import "PROJECT-III/domain"

type OrderFormat struct {
	ProductID int `json:"productid" form:"productid" validate:"required"`
	Qty       int `json:"qty" form:"qty" validate:"required"`
}

func (of *OrderFormat) ToOrder() domain.Order {
	return domain.Order{
		ProductID: of.ProductID,
		Qty:       of.Qty,
	}
}
