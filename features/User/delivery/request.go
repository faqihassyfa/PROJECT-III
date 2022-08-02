package delivery

import "PROJECT-III/domain"

type UserFormat struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
}

func (uf *UserFormat) ToModel() domain.User {
	return domain.User{
		Name:     uf.Name,
		Email:    uf.Email,
		Password: uf.Password,
		Address:  uf.Address,
		Phone:    uf.Phone,
	}
}

type LoginFormat struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (lf *LoginFormat) ToModelLogin() domain.User {
	return domain.User{
		Email:    lf.Email,
		Password: lf.Password,
	}
}

type ReturnFormat struct {
}
