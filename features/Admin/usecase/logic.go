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

func (auc *adminUsecase) CreateProduct(newProduct domain.Product, adminid int) int {
	var product = data.FromModel(newProduct)
	validError := auc.validate.Struct(product)

	if validError != nil {
		log.Println("Validation error : ", validError)
		return 400
	}

	product.Adminid = adminid
	create := auc.adminData.CreateProductData(product.ToModel())

	if create.ID == 0 {
		log.Println("error after creating data")
		return 500
	}
	return 200
}

func (auc *adminUsecase) ReadAllProduct(adminid int) ([]domain.Product, int) {
	products := auc.adminData.ReadAllProductData(adminid)
	if len(products) == 0 {
		log.Println("data not found")
		return nil, 404
	}
	return products, 200
}
