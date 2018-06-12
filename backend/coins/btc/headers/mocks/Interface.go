// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import headers "github.com/shiftdevices/godbb/backend/coins/btc/headers"
import mock "github.com/stretchr/testify/mock"
import wire "github.com/btcsuite/btcd/wire"

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// HeaderByHeight provides a mock function with given fields: _a0
func (_m *Interface) HeaderByHeight(_a0 int) (*wire.BlockHeader, error) {
	ret := _m.Called(_a0)

	var r0 *wire.BlockHeader
	if rf, ok := ret.Get(0).(func(int) *wire.BlockHeader); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wire.BlockHeader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields:
func (_m *Interface) Init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Status provides a mock function with given fields:
func (_m *Interface) Status() (*headers.Status, error) {
	ret := _m.Called()

	var r0 *headers.Status
	if rf, ok := ret.Get(0).(func() *headers.Status); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*headers.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeEvent provides a mock function with given fields: f
func (_m *Interface) SubscribeEvent(f func(headers.Event)) func() {
	ret := _m.Called(f)

	var r0 func()
	if rf, ok := ret.Get(0).(func(func(headers.Event)) func()); ok {
		r0 = rf(f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func())
		}
	}

	return r0
}

// TipHeight provides a mock function with given fields:
func (_m *Interface) TipHeight() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
