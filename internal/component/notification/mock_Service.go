// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package notification

import mock "github.com/stretchr/testify/mock"

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// List provides a mock function with given fields: customerID
func (_m *MockService) List(customerID int) ([]Notification, error) {
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

// Retry provides a mock function with given fields: customerID, notificationID
func (_m *MockService) Retry(customerID int, notificationID int) (int, error) {
	ret := _m.Called(customerID, notificationID)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(customerID, notificationID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(customerID, notificationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Send provides a mock function with given fields: notif
func (_m *MockService) Send(notif *Notification) (int, error) {
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

// SendInternal provides a mock function with given fields: typeID, custID, payload
func (_m *MockService) SendInternal(typeID int, custID int, payload interface{}) (int, error) {
	ret := _m.Called(typeID, custID, payload)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int, interface{}) int); ok {
		r0 = rf(typeID, custID, payload)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, interface{}) error); ok {
		r1 = rf(typeID, custID, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartLiseners provides a mock function with given fields:
func (_m *MockService) StartLiseners() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Test provides a mock function with given fields: webhookID
func (_m *MockService) Test(webhookID int) (int, error) {
	ret := _m.Called(webhookID)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(webhookID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(webhookID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: notifID, status
func (_m *MockService) UpdateStatus(notifID int, status string) error {
	ret := _m.Called(notifID, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(notifID, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
