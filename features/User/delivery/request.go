package delivery

import "PROJECT-III/domain"

type UserFormat struct {
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Address string `json:"address" form:"address"`
	Phone   string `json:"phone" form:"phone"`
}

func (i *UserFormat) ToModel() domain.User {
	return domain.User{
		Name:    i.Name,
		Email:   i.Email,
		Address: i.Address,
		Phone:   i.Phone,
	}
}
