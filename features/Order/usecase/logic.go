package usecase

import (
	"PROJECT-III/domain"
	"PROJECT-III/features/Order/data"
	"log"

	"github.com/go-playground/validator/v10"
)

type orderUseCase struct {
	orderData domain.OrderData
	validate  *validator.Validate
}

func New(od domain.OrderData, v *validator.Validate) domain.OrderUseCase {
	return &orderUseCase{
		orderData: od,
		validate:  v,
	}
}

func (ouc *orderUseCase) CreateOrder(neworder domain.Order, userid int) int {
	var order = data.FromModel(neworder)
	validError := ouc.validate.Struct(order)

	if validError != nil {
		log.Println("validation error : ", validError)
		return 400
	}

	order.Userid = userid
	create := ouc.orderData.CreateOrderData(order.ToModel())
	if create.ID == 0 {
		log.Println("error after creating data")
		return 500
	}
	return 200

}
