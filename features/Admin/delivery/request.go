package delivery

import "PROJECT-III/domain"

type ProductFormat struct {
	Name  string `json:"name" form:"name" validate:"required"`
	Price int    `json:"price" form:"price" validate:"required"`
	Stock int    `json:"stock" form:"stock" validate:"required"`
	Image string `json:"image" form:"image" validate:"required"`
}

func (pf *ProductFormat) ToModel() domain.Product {
	return domain.Product{
		Name:  pf.Name,
		Price: pf.Price,
		Stock: pf.Stock,
		Image: pf.Image,
	}
}
