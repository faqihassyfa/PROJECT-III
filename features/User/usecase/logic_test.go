package usecase

import (
	"PROJECT-III/domain"
	"PROJECT-III/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
	repo := new(mocks.UserData)

	mockData := domain.User{Name: "vanili", Email: "vanili@vanili.com", Password: "123", Address: "malang", Phone: "012345", Role: "users"}

	emptyMockData := domain.User{ID: 0, Name: "", Email: "", Password: "", Address: "", Phone: "", Role: ""}

	returnData := mockData
	returnData.ID = 1
	returnData.Password = "$2a$10$0f2Y2o5ptsvsT1r5GY87pOXQTL3gEATCn902I1RetP8g5doVbRgPW"

	invalidData := mockData
	invalidData.Email = ""

	noData := mockData
	noData.ID = 0

	t.Run("Success register data", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("RegisterData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(mockData)

		assert.Equal(t, 200, res)
		repo.AssertExpectations(t)
	})
	t.Run("Validation Error", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("RegisterData", mock.Anything).Return(invalidData).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(mockData)

		assert.Equal(t, 404, res)
		repo.AssertExpectations(t)
	})
	t.Run("Generate Hash Error", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("RegisterData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(mockData)

		assert.NotEqual(t, 500, res)
		repo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("RegisterData", mock.Anything).Return(emptyMockData).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(noData)

		assert.Equal(t, 404, res)
		repo.AssertExpectations(t)
	})
	t.Run("Duplicate Data", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(true).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(mockData)

		assert.Equal(t, 400, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	repo := new(mocks.UserData)

	mockData := domain.User{ID: 1, Name: "vanili", Email: "vanili@vanili.com", Password: "123", Address: "malang", Phone: "012345", Role: "users"}
	returnData := domain.User{ID: 1, Name: "vanili", Email: "vanili@vanili.com", Password: "123", Address: "malang", Phone: "012345", Role: "users"}

	invalidData := mockData
	invalidData.Name = ""

	t.Run("Success Update", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("UpdateUserData", mock.Anything, mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.UpdateUser(mockData, 1)

		assert.Equal(t, 200, res)
		repo.AssertExpectations(t)
	})
	t.Run("Validation Error", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("RegisterData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(invalidData)

		assert.NotEqual(t, 500, res)
		repo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res := useCase.UpdateUser(mockData, 0)

		assert.Equal(t, 404, res)
		repo.AssertExpectations(t)
	})
	t.Run("Generate Hash Error", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("UpdateUserData", mock.Anything, mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.UpdateUser(mockData, 1)

		assert.NotEqual(t, 500, res)
		repo.AssertExpectations(t)
	})
	t.Run("Duplicate Data", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(true).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(mockData)

		assert.Equal(t, 400, res)
		repo.AssertExpectations(t)
	})
}

func TestLoginUser(t *testing.T) {
	repo := new(mocks.UserData)
	mockData := domain.User{Email: "vanili@vanili.com", Password: "123"}
	returnData := domain.User{ID: 1}
	token := "abababababa122333"
	t.Run("Success Login", func(t *testing.T) {
		repo.On("GetPasswordData", mock.Anything).Return("$2a$10$0f2Y2o5ptsvsT1r5GY87pOXQTL3gEATCn902I1RetP8g5doVbRgPW")
		repo.On("LoginData", mock.Anything).Return(returnData, token).Once()
		userUseCase := New(repo, validator.New())
		res, err := userUseCase.LoginUser(mockData)

		assert.NotNil(t, err)
		assert.GreaterOrEqual(t, res.ID, 0)
		repo.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	repo := new(mocks.UserData)

	t.Run("Success Delete", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything).Return(true).Once()
		useCase := New(repo, validator.New())
		delete := useCase.DeleteUser(1)

		assert.Equal(t, 200, delete)
		repo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything).Return(true).Once()
		useCase := New(repo, validator.New())
		delete := useCase.DeleteUser(0)

		assert.NotEqual(t, 404, delete)
		repo.AssertExpectations(t)
	})
}

// func TestAccountUser(t *testing.T) {
// 	repo := new(mocks.UserData)

// 	returnDataUser := domain.User{ID: 1, Name: "vanili", Email: "vanili@vanili.com", Password: "123", Address: "malang", Phone: "012345", Role: "users"}
// 	returnDataOrderHistory := []domain.OrderHistory{{ID: 1, Totalprice: 150000}}

// 	t.Run("Success Get User", func(t *testing.T) {
// 		repo.On("AccountUserData", mock.Anything).Return(returnDataUser).Once()
// 		repo.On("HistoryUserData", mock.Anything).Return(returnDataOrderHistory)
// 		useCase := New(repo, validator.New())
// 		account, _, Order := useCase.AccountUser(1)

// 		assert.Equal(t, returnDataUser, account)
// 		assert.Equal(t, 200, Order)
// 		repo.AssertExpectations(t)
// 	})
// 	t.Run("Data not Found", func(t *testing.T) {
// 		returnDataUser.ID = 0
// 		repo.On("AccountUserData", mock.Anything).Return(returnDataUser).Once()
// 		repo.On("HistoryUserData", mock.Anything).Return(returnDataOrderHistory)
// 		useCase := New(repo, validator.New())
// 		account, _, order := useCase.AccountUser(1)

// 		assert.Equal(t, 0, account.ID)
// 		assert.Equal(t, 404, order)
// 		repo.AssertExpectations(t)
// 	})
// }

func TestAllProduct(t *testing.T) {
	repo := new(mocks.UserData)
	mockData := []domain.Product{{ID: 1, Adminid: 1, Name: "Tshirt", Price: 75000, Stock: 5, Image: "image.jpg"}}

	t.Run("Success Get Product", func(t *testing.T) {
		repo.On("AllProductData").Return(mockData).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.AllProduct()

		assert.Equal(t, 200, status)
		assert.GreaterOrEqual(t, len(res), 1)
		assert.Greater(t, res[0].ID, 0)
		repo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("AllProductData").Return([]domain.Product{}).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.AllProduct()

		assert.Equal(t, 404, status)
		assert.Equal(t, len(res), 0)
		assert.Equal(t, []domain.Product([]domain.Product(nil)), res)
		assert.Equal(t, []domain.Product(nil), res)
		repo.AssertExpectations(t)
	})
}
