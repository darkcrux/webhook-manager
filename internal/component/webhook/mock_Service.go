// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package webhook

import mock "github.com/stretchr/testify/mock"

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// Create provides a mock function with given fields: wh
func (_m *MockService) Create(wh *Webhook) (int, error) {
	ret := _m.Called(wh)

	var r0 int
	if rf, ok := ret.Get(0).(func(*Webhook) int); ok {
		r0 = rf(wh)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Webhook) error); ok {
		r1 = rf(wh)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *MockService) Get(id int) (*Webhook, error) {
	ret := _m.Called(id)

	var r0 *Webhook
	if rf, ok := ret.Get(0).(func(int) *Webhook); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Webhook)
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

// GetByTxAndCust provides a mock function with given fields: txId, custID
func (_m *MockService) GetByTxAndCust(txId int, custID int) (*Webhook, error) {
	ret := _m.Called(txId, custID)

	var r0 *Webhook
	if rf, ok := ret.Get(0).(func(int, int) *Webhook); ok {
		r0 = rf(txId, custID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Webhook)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(txId, custID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: id
func (_m *MockService) List(id int) ([]Webhook, error) {
	ret := _m.Called(id)

	var r0 []Webhook
	if rf, ok := ret.Get(0).(func(int) []Webhook); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Webhook)
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

// Update provides a mock function with given fields: customerID, webhookID, url
func (_m *MockService) Update(customerID int, webhookID int, url string) (int, error) {
	ret := _m.Called(customerID, webhookID, url)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int, string) int); ok {
		r0 = rf(customerID, webhookID, url)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(customerID, webhookID, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
