package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ID        int
	Adminid   int
	Name      string
	Price     int
	Stock     int
	Image     string
	CreatedAt time.Time
	Orders    []Order
}

type AdminOrderHistory struct {
	ID         int
	CreatedAt  time.Time
	Totalprice int
}

type AdminHandler interface {
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Create() echo.HandlerFunc
	ReadAll() echo.HandlerFunc
	// ReadHistory() echo.HandlerFunc
}

type AdminUseCase interface {
	UpdateProduct(updatedData Product, productid, adminid int) int
	DeleteProduct(productid, adminid int) int
	CreateProduct(newProduct Product, adminid int) int
	ReadAllProduct(adminid int) ([]Product, int)
	// HistoryAdmin(adminid int) ([]AdminOrderHistory, int)
}

type AdminData interface {
	UpdateProductData(updatedData Product) Product
	DeleteProductData(productid, adminid int) bool
	CreateProductData(newProduct Product) Product
	ReadAllProductData(adminid int) []Product
	// HistoryAdminData(adminid int) []AdminOrderHistory
}
