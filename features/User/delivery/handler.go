package delivery

import (
	"PROJECT-III/domain"
	"PROJECT-III/features/common"
	"log"
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
		var newuser UserFormat
		bind := c.Bind(&newuser)

		if bind != nil {
			log.Println("cannot bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in internal sever",
			})
		}

		status := uh.userUserCase.RegisterUser(newuser.ToModel())

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "Wrong input",
			})
		}
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
			"message": "Register Success",
		})
	}
}

func (uh *userHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp UserFormat
		res := c.Bind(&tmp)

		qry := map[string]interface{}{}
		id := common.ExtractData(c)

		if res != nil {
			log.Println("Cannot parse data", res)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		if tmp.Name != "" {
			qry["name"] = tmp.Name
		}

		if tmp.Email != "" {
			qry["email"] = tmp.Email
		}

		if tmp.Address != "" {
			qry["address"] = tmp.Address
		}

		if tmp.Phone != "" {
			qry["phone"] = tmp.Phone
		}

		status := uh.userUserCase.UpdateUser(tmp.ToModel(), id)

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "wrong input",
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
			"message": "Registrasi Succes",
		})
	}
}

func (uh *userHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)
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

func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var datalogin LoginFormat
		bind := c.Bind(&datalogin)

		if bind != nil {
			log.Println("invalid input")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in internal server",
			})
		}

		data, err := uh.userUserCase.LoginUser(datalogin.ToModelLogin())

		if err != nil {
			log.Println("Login failed", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "Wrong username or password",
			})
		}

		token := common.GenerateToken(int(data.ID))

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "Login success",
			"token":   token,
			"role":    data.Role,
		})
	}
}
