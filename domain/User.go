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
	Orders   []Order
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

type UserHandler interface {
	Register() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Login() echo.HandlerFunc
	// Account() echo.HandlerFunc
	Product() echo.HandlerFunc
}

type UserUseCase interface {
	RegisterUser(newuser User) int
	DeleteUser(userID int) int
	LoginUser(userdata User) (User, error)
	// AccountUser(userid int) (User, []OrderHistory, int)
	UpdateUser(updatedData User, userid int) int
	AllProduct() ([]Product, int)
}

type UserData interface {
	RegisterData(newuser User) User
	DeleteData(userID int) bool
	LoginData(userdata User) User
	GetPasswordData(name string) string
	AccountUserData(userid int) User
	// HistoryUserData(userid int) []OrderHistory
	UpdateUserData(updatedData User, userid int) User
	CheckDuplicate(newuser User) bool
	AllProductData() []Product
}
