// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package customer

import mock "github.com/stretchr/testify/mock"

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: id
func (_m *MockRepository) GetByID(id int) (*Customer, error) {
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

// Save provides a mock function with given fields: c
func (_m *MockRepository) Save(c *Customer) (int, error) {
	ret := _m.Called(c)

	var r0 int
	if rf, ok := ret.Get(0).(func(*Customer) int); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Customer) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
