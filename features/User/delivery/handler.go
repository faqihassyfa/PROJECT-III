package delivery

import (
	"PROJECT-III/domain"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUserCase domain.UserUseCase
	userData     domain.UserData
}

func New(uuc domain.UserUseCase, ud domain.UserData) domain.UserHandler {
	return &userHandler{
		userUserCase: uuc,
		userData:     ud,
	}
}

// Register implementasi domain.UserHandler
func (uh *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newuser UserFormat
		bind := c.Bind(&newuser)

		if bind != nil {
			log.Println("cannot bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in the internal server",
			})
		}
		status := uh.userUserCase.RegisterUser(newuser.ToModel())

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "Wrong input",
			})
		}

		if status == 400 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "Data Not Found",
			})
		}
		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "There is an error in the internal server",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": "Register Success",
		})
	}
}

func (uh *userHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtraData(c)
		status := uh.userUserCase.DeleteUser(id)

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "Data Not Found",
			})
		}
		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "There is an error in the internal server",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": "Success Delete Data",
		})
	}
}
