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

type OrderHistory struct {
	ID        int
	CreatedAt time.Time
	Total     int
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Login() echo.HandlerFunc
	Account() echo.HandlerFunc
}

type UserUseCase interface {
	RegisterUser(newuser User, IDuser int) int
	LoginUser(authData LoginAuth) (data map[string]interface{}, err error)
	AccountUser(userid int) (User, []OrderHistory, int)
}

type UserData interface {
	RegisterData(newuser User) User
	LoginData(authData LoginAuth) (data map[string]interface{}, err error)
	AccountUserData(userid int) User
	HistoryUserData(userid int) []OrderHistory
}
