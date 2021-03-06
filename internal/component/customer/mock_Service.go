// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package customer

import mock "github.com/stretchr/testify/mock"

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// Get provides a mock function with given fields: id
func (_m *MockService) Get(id int) (*Customer, error) {
	ret := _m.Called(id)

	var r0 *Customer
	if rf, ok := ret.Get(0).(func(int) *Customer); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: tx
func (_m *MockService) Register(tx *Customer) (int, error) {
	ret := _m.Called(tx)

	var r0 int
	if rf, ok := ret.Get(0).(func(*Customer) int); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Customer) error); ok {
		r1 = rf(tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
