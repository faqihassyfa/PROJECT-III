package usecase

// func TestCreateOrder(t *testing.T) {
// 	repo := new(mocks.OrderData)

// 	mockData := domain.Order{ID: 1, Userid: 1, PaymentMethod: "bank", Qty: 2, Totalprice: 150000, Status: 1, Productid: 1}
// 	emptyMockData := domain.Order{ID: 0, PaymentMethod: "", Qty: 0, Totalprice: 0, Status: 0, Productid: 0}

// 	returnData := mockData
// 	returnData.ID = 1

// 	invalidData := mockData
// 	invalidData.PaymentMethod = ""

// 	noData := mockData
// 	noData.ID = 0

// 	t.Run("Success Input Data", func(t *testing.T) {
// 		repo.On("CreateOrderData", mock.Anything).Return(returnData).Once()
// 		useCase := New(repo, validator.New())
// 		res := useCase.CreateOrder(mockData, 1)

// 		assert.Equal(t, 200, res)
// 		repo.AssertExpectations(t)
// 	})
// 	t.Run("Valiadtion Error", func(t *testing.T) {
// 		repo.On("CreateOrderData", mock.Anything).Return(returnData).Once()
// 		useCase := New(repo, validator.New())
// 		res := useCase.CreateOrder(invalidData, 1)

// 		assert.NotEqual(t, 400, res)
// 		repo.AssertExpectations(t)
// 	})
// 	t.Run("Error after creating data", func(t *testing.T) {
// 		repo.On("CreateOrderData", mock.Anything).Return(returnData).Once()
// 		useCase := New(repo, validator.New())
// 		res := useCase.CreateOrder(mockData, 1)

// 		assert.NotEqual(t, 500, res)
// 		repo.AssertExpectations(t)
// 	})
// 	t.Run("Data Not Found", func(t *testing.T) {
// 		repo.On("CreateOrderData", mock.Anything).Return(returnData).Once()
// 		useCase := New(repo, validator.New())
// 		res := useCase.CreateOrder(emptyMockData, 0)

// 		assert.NotEqual(t, 404, res)
// 		repo.AssertExpectations(t)
// 	})
// }
