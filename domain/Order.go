package domain

import "github.com/labstack/echo"

type Order struct {
	ID            int
	Userid        int
	PaymentMethod string
	TotalQty      int
	Totalprice    int
	Status        int
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
