package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Order struct {
	ID         int
	OrderID    string
	UserID     int
	ProductID  int
	Price      int
	Qty        int
	Totalprice int
	Status     int
	Payment    string
	CreatedAt  time.Time
}

type OrderHandler interface {
	Create() echo.HandlerFunc
}

type OrderUseCase interface {
	CreateOrder(neworder Order, userid int) (int, string)
	ConfirmStatus(orderID int) int
	CancelStatus(orderID int) int
}

type OrderData interface {
	InsertOrderData(neworder Order) Order
	TotalPriceOrder(neworder Order) (int, int)
	CekStock(neworder Order) bool
	ReturnStock(neworder Order)
	ConfirmStatusData(idOrder int) (row int)
	CancelStatusData(orderID int) (row int)
}
