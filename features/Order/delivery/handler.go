package delivery

import (
	"PROJECT-III/domain"
	"PROJECT-III/features/common"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type orderHandler struct {
	orderData    domain.OrderData
	orderUseCase domain.OrderUseCase
}

func New(od domain.OrderData, ouc domain.OrderUseCase) domain.OrderHandler {
	return &orderHandler{
		orderUseCase: ouc,
		orderData:    od,
	}
}

func (oh *orderHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var neworder OrderFormat
		id, _ := common.ExtractData(c)
		bind := c.Bind(&neworder)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in internal server",
			})
		}

		status := oh.orderUseCase.CreateOrder(neworder.ToModel(), id)
		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "Wrong input",
			})
		}

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "Data not found",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "There is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": "Order success",
		})
	}
}
