package delivery

import (
	"PROJECT-III/domain"
	"PROJECT-III/features/common"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUserCase domain.UserUseCase
	userdata     domain.UserData
}

func New(uuc domain.UserUseCase, ud domain.UserData) domain.UserHandler {
	return &userHandler{
		userUserCase: uuc,
		userdata:     ud,
	}
}

func (uh *userHandler) Account() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)
		var arrmap []map[string]interface{}

		myaccount, myorder, status := uh.userUserCase.AccountUser(id)

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "Data not found",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "There is an error in internal server",
			})
		}

		var res = map[string]interface{}{}
		res["name"] = myaccount.Name
		res["email"] = myaccount.Email
		res["address"] = myaccount.Address
		res["phone"] = myaccount.Phone

		arrmap = append(arrmap, res)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"my account":       res,
			"my order history": myorder,
			"code":             200,
			"message":          "Succes Get My Account",
		})
	}
}

func (uh *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "Registrasi Succes",
		})
	}
}

func (uh *userHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "Registrasi Succes",
		})
	}
}

func (uh *userHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "Registrasi Succes",
		})
	}
}

func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "Registrasi Succes",
		})
	}
}
