package data

import (
	"PROJECT-III/domain"

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

type UserProfile struct {
	ID    int
	Name  string
	Phone string
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
