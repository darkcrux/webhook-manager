// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package transport

import mock "github.com/stretchr/testify/mock"

// MockMessageBus is an autogenerated mock type for the MessageBus type
type MockMessageBus struct {
	mock.Mock
}

// PublishNotification provides a mock function with given fields: notif
func (_m *MockMessageBus) PublishNotification(notif *Notification) error {
	ret := _m.Called(notif)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Notification) error); ok {
		r0 = rf(notif)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublishNotificationStatus provides a mock function with given fields: notifStatus
func (_m *MockMessageBus) PublishNotificationStatus(notifStatus *NotificationStatus) error {
	ret := _m.Called(notifStatus)

	var r0 error
	if rf, ok := ret.Get(0).(func(*NotificationStatus) error); ok {
		r0 = rf(notifStatus)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscribeNotificationStatus provides a mock function with given fields: _a0
func (_m *MockMessageBus) SubscribeNotificationStatus(_a0 func(*NotificationStatus) error) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(*NotificationStatus) error) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubsribeNotification provides a mock function with given fields: _a0
func (_m *MockMessageBus) SubsribeNotification(_a0 func(*Notification) error) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(*Notification) error) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}