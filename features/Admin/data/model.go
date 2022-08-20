package data

import (
	"PROJECT-III/domain"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Adminid int
	Name    string `json:"name" form:"name" validate:"required"`
	Price   int    `json:"price" form:"price" validate:"required"`
	Stock   int    `json:"stock" form:"stock" validate:"required"`
	Image   string `json:"image" form:"image"`
	// Orders  []data.Orders `gorm:"foreignKey:Productid"`
}

type AdminOrderHistory struct {
	ID         int
	CreatedAt  time.Time
	Totalprice int
}

func (u *Product) ToModel() domain.Product {
	return domain.Product{
		ID:        int(u.ID),
		Adminid:   u.Adminid,
		Name:      u.Name,
		Price:     u.Price,
		Stock:     u.Stock,
		Image:     u.Image,
		CreatedAt: u.CreatedAt,
	}
}
func (p *AdminOrderHistory) ToAdminOrderHistory() domain.AdminOrderHistory {
	return domain.AdminOrderHistory{
		ID:         p.ID,
		CreatedAt:  p.CreatedAt,
		Totalprice: p.Totalprice,
	}
}

func ParseToArr(arr []Product) []domain.Product {
	var res []domain.Product

	for _, val := range arr {
		res = append(res, val.ToModel())
	}

	return res
}

func ParseAdminOrderHistoryToArr(arr []AdminOrderHistory) []domain.AdminOrderHistory {
	var res []domain.AdminOrderHistory

	for _, val := range arr {
		res = append(res, val.ToAdminOrderHistory())
	}

	return res
}

func FromModel(data domain.Product) Product {
	var res Product
	res.ID = uint(data.ID)
	res.Adminid = data.Adminid
	res.Name = data.Name
	res.Price = data.Price
	res.Stock = data.Stock
	res.Image = data.Image
	res.CreatedAt = data.CreatedAt
	return res
}
