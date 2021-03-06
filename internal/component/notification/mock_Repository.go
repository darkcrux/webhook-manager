// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package notification

import mock "github.com/stretchr/testify/mock"

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: notif
func (_m *MockRepository) Create(notif *Notification) (int, error) {
	ret := _m.Called(notif)

	var r0 int
	if rf, ok := ret.Get(0).(func(*Notification) int); ok {
		r0 = rf(notif)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Notification) error); ok {
		r1 = rf(notif)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *MockRepository) Get(id int) (*Notification, error) {
	ret := _m.Called(id)

	var r0 *Notification
	if rf, ok := ret.Get(0).(func(int) *Notification); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Notification)
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

// List provides a mock function with given fields: customerID
func (_m *MockRepository) List(customerID int) ([]Notification, error) {
	ret := _m.Called(customerID)

	var r0 []Notification
	if rf, ok := ret.Get(0).(func(int) []Notification); ok {
		r0 = rf(customerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Notification)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(customerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: id, status
func (_m *MockRepository) UpdateStatus(id int, status string) error {
	ret := _m.Called(id, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(id, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
