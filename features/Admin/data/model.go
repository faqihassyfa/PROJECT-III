package data

import (
	"PROJECT-III/domain"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Adminid int
	Name    string `json:"name" form:"name" validate:"required"`
	Price   int    `json:"price" form:"price" validate:"required"`
	Stock   int    `json:"stock" form:"stock" validate:"required"`
	Image   string `json:"image" form:"image"`
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

func ParseToArr(arr []Product) []domain.Product {
	var res []domain.Product

	for _, val := range arr {
		res = append(res, val.ToModel())
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
