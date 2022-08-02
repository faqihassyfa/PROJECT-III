package delivery

import "PROJECT-III/domain"

type UserFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Phone    string `json:"phone" form:"phone"`
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
