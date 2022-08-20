package usecase

import (
	"PROJECT-III/domain"
	"PROJECT-III/infrastructure/database/midtrans"
	"log"

	"github.com/midtrans/midtrans-go/snap"
)

type orderUseCase struct {
	orderData domain.OrderData
	midtrans  snap.Client
}

func New(od domain.OrderData, mt snap.Client) domain.OrderUseCase {
	return &orderUseCase{
		orderData: od,
		midtrans:  mt,
	}
}

func (ouc *orderUseCase) CreateOrder(neworder domain.Order, userid int) (int, string) {

	neworder.UserID = userid
	orderprice, ordertotal := ouc.orderData.TotalPriceOrder(neworder)
	neworder.Price = orderprice
	neworder.Totalprice = ordertotal
	stock := ouc.orderData.CekStock(neworder)

	if !stock {
		log.Println("not enough stock")
		return 404, ""
	}

	create := ouc.orderData.InsertOrderData(neworder)
	if create.ID == 0 {
		log.Println("error after creating data")
		return 404, ""
	}

	res, orderid := midtrans.Payment(neworder.Totalprice)
	neworder.Payment = res.RedirectURL
	neworder.Status = 0
	neworder.OrderID = orderid

	return 200, neworder.Payment
}

func (ouc *orderUseCase) ConfirmStatus(orderID int) int {
	row := ouc.orderData.ConfirmStatusData(orderID)
	return row
}

func (ouc *orderUseCase) CancelStatus(orderID int) int {
	row := ouc.orderData.CancelStatusData(orderID)
	return row
}
