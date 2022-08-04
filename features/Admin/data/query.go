package data

import (
	"PROJECT-III/domain"
	dataorder "PROJECT-III/features/Order/data"
	"log"

	"gorm.io/gorm"
)

type adminData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.AdminData {
	return &adminData{
		db: db,
	}
}

func (ad *adminData) UpdateProductData(updatedData domain.Product) domain.Product {
	var detailuser domain.User
	var products = FromModel(updatedData)
	cekadmin := ad.db.Table("users").Where("ID = ?", products.Adminid).First(&detailuser)

	if cekadmin.Error != nil {
		log.Println("cannot read data", cekadmin.Error.Error())
		return domain.Product{}
	}

	if detailuser.Role != "admin" {
		log.Println("not admin!", cekadmin.Error.Error())
		return domain.Product{}
	}

	err := ad.db.Model(&Product{}).Where("ID = ?", products.ID).Updates(products)

	if err.Error != nil {
		log.Println("cannot update data", err.Error.Error())
		return domain.Product{}
	}

	if err.RowsAffected == 0 {
		log.Println("data not found")
		return domain.Product{}
	}
	return products.ToModel()
}

func (ad *adminData) DeleteProductData(productid, adminid int) bool {
	var detailuser domain.User
	cekadmin := ad.db.Table("users").Where("ID = ?", adminid).First(&detailuser)

	if cekadmin.Error != nil {
		log.Println("cannot read data", cekadmin.Error.Error())
		return false
	}

	if detailuser.Role != "admin" {
		log.Println("not admin!", cekadmin.Error.Error())
		return false
	}

	err := ad.db.Where("ID = ?", productid).Delete(&Product{})

	if err.Error != nil {
		log.Println("Cannot delete data", err.Error.Error())
		return false
	}
	if err.RowsAffected < 1 {
		log.Println("No data deleted", err.Error.Error())
		return false
	}
	return true
}

// func (ad *adminData) CreateProductData(newProduct domain.Product) domain.Product {
// 	var product = FromModel(newProduct)
// 	err := ad.db.Create(&product)

// 	if err.Error != nil {
// 		log.Println("cannot create data", err.Error)
// 		return domain.Product{}
// 	}
// 	return product.ToModel()
// }

func (ad *adminData) CreateProductData(newProduct domain.Product) domain.Product {
	var detailuser domain.User
	var products = FromModel(newProduct)
	cekadmin := ad.db.Table("users").Where("ID = ?", products.Adminid).First(&detailuser)

	if cekadmin.Error != nil {
		log.Println("cannot read data", cekadmin.Error.Error())
		return domain.Product{}
	}

	if detailuser.Role != "admin" {
		log.Println("not admin!", cekadmin.Error.Error())
		return domain.Product{}
	}

	err := ad.db.Create(&products)

	if err.Error != nil {
		log.Println("cannot create data", err.Error.Error())
		return domain.Product{}
	}

	if err.RowsAffected == 0 {
		log.Println("data not found")
		return domain.Product{}
	}
	return products.ToModel()
}

func (ad *adminData) ReadAllProductData(adminid int) []domain.Product {
	var detailuser domain.User
	var products []Product
	cekadmin := ad.db.Table("users").Where("ID = ?", adminid).First(&detailuser)

	if cekadmin.Error != nil {
		log.Println("cannot read data", cekadmin.Error.Error())
		return []domain.Product{}
	}

	if detailuser.Role != "admin" {
		log.Println("not admin!", cekadmin.Error.Error())
		return []domain.Product{}
	}
	err := ad.db.Find(&products)
	if err.Error != nil {
		log.Println("cannot read data", err.Error.Error())
		return []domain.Product{}
	}
	if err.RowsAffected == 0 {
		log.Println("data not found")
		return []domain.Product{}
	}
	return ParseToArr(products)
}

func (ad *adminData) HistoryAdminData(adminid int) []domain.AdminOrderHistory {
	var detailuser domain.User
	var orders []AdminOrderHistory
	cekadmin := ad.db.Table("users").Where("ID = ?", adminid).First(&detailuser)

	if cekadmin.Error != nil {
		log.Println("cannot read data", cekadmin.Error.Error())
		return []domain.AdminOrderHistory{}
	}

	if detailuser.Role != "admin" {
		log.Println("not admin!", cekadmin.Error.Error())
		return []domain.AdminOrderHistory{}
	}
	err := ad.db.Model(&dataorder.Order{}).Select("orders.totalprice, orders.created_at, orders.id").Where("userid = ?", adminid).Find(&orders).Error
	if err != nil {
		log.Println("cannot read data", err.Error())
		return []domain.AdminOrderHistory{}
	}

	return ParseAdminOrderHistoryToArr(orders)
}
