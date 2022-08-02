package delivery

import (
	"PROJECT-III/domain"
<<<<<<< HEAD
=======
	"PROJECT-III/features/common"
>>>>>>> d93fae70d318a208749989d806637816f93eb39c
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUserCase domain.UserUseCase
<<<<<<< HEAD
	userData     domain.UserData
=======
	userdata     domain.UserData
>>>>>>> d93fae70d318a208749989d806637816f93eb39c
}

func New(uuc domain.UserUseCase, ud domain.UserData) domain.UserHandler {
	return &userHandler{
		userUserCase: uuc,
<<<<<<< HEAD
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
=======
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
>>>>>>> d93fae70d318a208749989d806637816f93eb39c

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
<<<<<<< HEAD
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
=======
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
>>>>>>> d93fae70d318a208749989d806637816f93eb39c
		})
	}
}

func (uh *userHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
<<<<<<< HEAD
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
=======
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
>>>>>>> d93fae70d318a208749989d806637816f93eb39c
		})
	}
}
