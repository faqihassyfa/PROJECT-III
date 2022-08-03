package usecase

import (
	"PROJECT-III/domain"
	"PROJECT-III/features/Admin/data"
	"log"

	"github.com/go-playground/validator/v10"
)

type adminUsecase struct {
	adminData domain.AdminData
	validate  *validator.Validate
}

func New(ad domain.AdminData, v *validator.Validate) domain.AdminUseCase {
	return &adminUsecase{
		adminData: ad,
		validate:  v,
	}
}

func (auc *adminUsecase) UpdateProduct(updatedData domain.Product, productid, adminid int) int {
	var products = data.FromModel(updatedData)
	products.ID = uint(productid)
	products.Adminid = adminid

	if productid == 0 {
		log.Println("Data not found")
		return 404
	}

	update := auc.adminData.UpdateProductData(products.ToModel())

	if update.ID == 0 {
		log.Println("empty data")
		return 500
	}
	return 200
}

func (auc *adminUsecase) DeleteProduct(productid, adminid int) int {
	status := auc.adminData.DeleteProductData(productid, adminid)

	if !status {
		log.Println("data not found")
		return 404
	}

	return 200
}
