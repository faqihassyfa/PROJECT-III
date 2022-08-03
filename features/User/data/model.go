package data

import (
	"PROJECT-III/domain"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Phone    string `json:"phone" form:"phone"`
	Role     string `json:"role" form:"role"`
}

type OrderHistory struct {
	ID         int
	CreatedAt  time.Time
	Totalprice int
}

type Myaccount struct {
	ID      int
	Name    string
	Email   string
	Address string
	Phone   string
}

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
	Image string
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:       int(u.ID),
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Address:  u.Address,
		Phone:    u.Phone,
		Role:     u.Role,
	}
}

func (u *Myaccount) ToMyAccount() domain.Myaccount {
	return domain.Myaccount{
		ID:      u.ID,
		Name:    u.Name,
		Email:   u.Email,
		Address: u.Address,
		Phone:   u.Phone,
	}
}

func (p *OrderHistory) ToOrderHistory() domain.OrderHistory {
	return domain.OrderHistory{
		ID:         p.ID,
		CreatedAt:  p.CreatedAt,
		Totalprice: p.Totalprice,
	}
}

func FromMyAccount(data domain.Myaccount) Myaccount {
	var res Myaccount
	res.ID = data.ID
	res.Name = data.Name
	res.Email = data.Email
	res.Address = data.Address
	res.Phone = data.Phone

	return res
}

func FromModel(data domain.User) User {
	var res User
	res.ID = uint(data.ID)
	res.Name = data.Name
	res.Email = data.Email
	res.Password = data.Password
	res.Address = data.Address
	res.Phone = data.Phone
	res.Role = data.Role

	return res
}

func ParseOrderHistoryToArr(arr []OrderHistory) []domain.OrderHistory {
	var res []domain.OrderHistory

	for _, val := range arr {
		res = append(res, val.ToOrderHistory())
	}

	return res
}

func (u *Product) ToProduct() domain.Product {
	return domain.Product{
		ID:    u.ID,
		Name:  u.Name,
		Price: u.Price,
		Stock: u.Stock,
		Image: u.Image,
	}
}

func ParseProductToArr(arr []Product) []domain.Product {
	var res []domain.Product

	for _, val := range arr {
		res = append(res, val.ToProduct())
	}

	return res
}
