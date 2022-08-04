package usecase

import (
	"PROJECT-III/domain"
	"PROJECT-III/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateProduct(t *testing.T) {
	repo := new(mocks.AdminData)

	mockData := domain.Product{Adminid: 1, Name: "Tshirt", Price: 75000, Stock: 5, Image: "image.jpg"}
	returnData := domain.Product{ID: 1, Adminid: 1, Name: "Tshirt", Price: 75000, Stock: 5, Image: "image.jpg"}

	invalidData := mockData
	invalidData.Name = ""

	t.Run("Success Update Product", func(t *testing.T) {
		repo.On("UpdateProductData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.UpdateProduct(mockData, 1, 1)

		assert.Equal(t, 200, res)
		repo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res := useCase.UpdateProduct(mockData, 0, 1)

		assert.Equal(t, 404, res)
		repo.AssertExpectations(t)
	})
}

func TestDeleteProduct(t *testing.T) {
	repo := new(mocks.AdminData)

	t.Run("Success Delete Product", func(t *testing.T) {
		repo.On("DeleteProductData", mock.Anything, mock.Anything).Return(true).Once()
		useCase := New(repo, validator.New())
		delete := useCase.DeleteProduct(0, 1) // masalah disini provided 1 argument, mocked 2 arguments

		assert.Equal(t, 200, delete)
		repo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("DeleteProductData", mock.Anything, mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		delete := useCase.DeleteProduct(1, 1)

		assert.Equal(t, 404, delete)
		repo.AssertExpectations(t)
	})
}

func TestCreateProduct(t *testing.T) {
	repo := new(mocks.AdminData)

	mockData := domain.Product{Adminid: 1, Name: "Tshirt", Price: 75000, Stock: 5, Image: "image.jpg"}
	emptyMockData := domain.Product{ID: 0, Adminid: 0, Name: "", Price: 0, Stock: 0, Image: ""}

	returnData := mockData
	returnData.ID = 1

	invalidData := mockData
	invalidData.Name = ""

	noData := mockData
	noData.ID = 0

	t.Run("Success insert product", func(t *testing.T) {
		repo.On("CreateProductData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.CreateProduct(mockData, 1)

		assert.Equal(t, 200, res)
		repo.AssertExpectations(t)
	})
	t.Run("Validation Error", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res := useCase.CreateProduct(invalidData, 1)

		assert.Equal(t, 400, res)
		repo.AssertExpectations(t)
	})
	// t.Run("error after creating data", func(t *testing.T) {
	// 	useCase := New(repo, validator.New())
	// 	res := useCase.CreateProduct(mockData, 1)

	// 	assert.Equal(t, 500, res)
	// 	repo.AssertExpectations(t)
	// })
	t.Run("Data Not Found", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res := useCase.CreateProduct(emptyMockData, 0)

		assert.NotEqual(t, 404, res)
		repo.AssertExpectations(t)
	})
}

func TestReadAllProduct(t *testing.T) {
	repo := new(mocks.AdminData)
	returnData := []domain.Product{{ID: 1, Adminid: 1, Name: "Tshirt", Price: 75000, Stock: 5, Image: "image.jpg"}}

	t.Run("Success Read Product", func(t *testing.T) {
		repo.On("ReadAllProductData").Return(returnData).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.ReadAllProduct()

		assert.Equal(t, 200, status)
		assert.GreaterOrEqual(t, len(res), 1)
		assert.Greater(t, res[0].ID, 0)
		repo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("ReadAllProductData").Return([]domain.Product{}).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.ReadAllProduct()

		assert.Equal(t, 404, status)
		assert.Equal(t, len(res), 0)
		assert.Equal(t, []domain.Product([]domain.Product(nil)), res)
		assert.Equal(t, []domain.Product(nil), res)
		repo.AssertExpectations(t)
	})
}
