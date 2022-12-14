// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "PROJECT-III/domain"

	mock "github.com/stretchr/testify/mock"
)

// OrderData is an autogenerated mock type for the OrderData type
type OrderData struct {
	mock.Mock
}

// CreateOrderData provides a mock function with given fields: neworder
func (_m *OrderData) CreateOrderData(neworder domain.Order) domain.Order {
	ret := _m.Called(neworder)

	var r0 domain.Order
	if rf, ok := ret.Get(0).(func(domain.Order) domain.Order); ok {
		r0 = rf(neworder)
	} else {
		r0 = ret.Get(0).(domain.Order)
	}

	return r0
}

type mockConstructorTestingTNewOrderData interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrderData creates a new instance of OrderData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrderData(t mockConstructorTestingTNewOrderData) *OrderData {
	mock := &OrderData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
