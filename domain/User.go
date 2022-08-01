package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Address  string
	Phone    string
	Role     string
}

type UserProfile struct {
	ID        int
	Name      string
	Phone     string
	CreatedAt time.Time
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Login() echo.HandlerFunc
	//GetProfile() echo.HandlerFunc
}

type UserUseCase interface {
	RegisterUser(newuser User, IDuser int) int
	LoginUser(authData LoginAuth) (data map[string]interface{}, err error)
}

type UserData interface {
	RegisterData(newuser User) User
	LoginData(authData LoginAuth) (data map[string]interface{}, err error)
}