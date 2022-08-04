package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Order struct {
	ID            int
	Userid        int
	PaymentMethod string
	Qty           int
	Totalprice    int
	Status        int
	Productid     int
	CreatedAt     time.Time
}

type ProductOrder struct {
	ID    int
	Name  string
	Price int
	Qty   int
}

type OrderHandler interface {
	Create() echo.HandlerFunc
}

type OrderUseCase interface {
	CreateOrder(neworder Order, userid int) int
}

type OrderData interface {
	CreateOrderData(neworder Order) Order
}
