package delivery

import (
	"PROJECT-III/domain"
	"PROJECT-III/features/common"
	awss3 "PROJECT-III/infrastructure/database/aws"
	"log"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo/v4"
)

type adminHandler struct {
	adminData    domain.AdminData
	adminUseCase domain.AdminUseCase
	conn         *session.Session
}

func New(ad domain.AdminData, auc domain.AdminUseCase, awscon *session.Session) domain.AdminHandler {
	return &adminHandler{
		adminUseCase: auc,
		adminData:    ad,
		conn:         awscon,
	}
}

func (ah *adminHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp ProductFormat
		bind := c.Bind(&tmp)

		qry := map[string]interface{}{}
		adminid := common.ExtractData(c)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "there is an error in internal server",
			})
		}

		productid := c.Param("productid")
		id, err := strconv.Atoi(productid)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "wrong input",
			})
		}

		if tmp.Name != "" {
			qry["name"] = tmp.Name
		}

		if tmp.Price != 0 {
			qry["price"] = tmp.Price
		}

		if tmp.Stock != 0 {
			qry["stock"] = tmp.Stock
		}

		file, err := c.FormFile("photo")

		if err != nil {
			log.Println(err)
		}

		link := awss3.DoUpload(ah.conn, *file, file.Filename)
		tmp.Image = link

		status := ah.adminUseCase.UpdateProduct(tmp.ToModel(), id, adminid)

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "wrong input",
			})
		}

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "data not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": "update success",
		})
	}
}

func (ah *adminHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		param := c.Param("productid")
		cnv, err := strconv.Atoi(param)

		if err != nil {
			log.Println("cant convert to int", err)
			return c.JSON(http.StatusInternalServerError, "cant convert to int")
		}

		id := common.ExtractData(c)

		status := ah.adminUseCase.DeleteProduct(cnv, id)

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
			"message": "Success delete product",
		})
	}
}

func (ah *adminHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newproduct ProductFormat
		id := common.ExtractData(c)
		bind := c.Bind(&newproduct)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "there is an error in internal server",
			})
		}

		file, err := c.FormFile("image")

		if err != nil {
			log.Println(err)
		}

		link := awss3.DoUpload(ah.conn, *file, file.Filename)
		newproduct.Image = link
		status := ah.adminUseCase.CreateProduct(newproduct.ToModel(), id)

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "wrong input",
			})
		}
		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "data not found",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "there is an error in internal server",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": "success create product",
		})
	}
}

func (ah *adminHandler) ReadAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, status := ah.adminUseCase.ReadAllProduct()

		var arrmap []map[string]interface{}
		var statuscode = map[string]interface{}{}
		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "data not found",
			})
		}
		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "there is an error in internal server",
			})
		}
		for i := 0; i < len(data); i++ {
			var res = map[string]interface{}{}

			res["id"] = data[i].ID
			res["name"] = data[i].Name
			res["stock"] = data[i].Stock
			res["price"] = data[i].Price

			arrmap = append(arrmap, res)
		}
		statuscode["code"] = status
		statuscode["messages"] = "success get all product"
		arrmap = append(arrmap, statuscode)

		return c.JSON(http.StatusOK, arrmap)
	}
}
