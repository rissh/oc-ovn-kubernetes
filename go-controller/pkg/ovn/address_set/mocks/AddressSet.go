// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	ovsdb "github.com/ovn-org/libovsdb/ovsdb"
	mock "github.com/stretchr/testify/mock"
)

// AddressSet is an autogenerated mock type for the AddressSet type
type AddressSet struct {
	mock.Mock
}

// AddAddresses provides a mock function with given fields: addresses
func (_m *AddressSet) AddAddresses(addresses []string) error {
	ret := _m.Called(addresses)

	if len(ret) == 0 {
		panic("no return value specified for AddAddresses")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(addresses)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddAddressesReturnOps provides a mock function with given fields: addresses
func (_m *AddressSet) AddAddressesReturnOps(addresses []string) ([]ovsdb.Operation, error) {
	ret := _m.Called(addresses)

	if len(ret) == 0 {
		panic("no return value specified for AddAddressesReturnOps")
	}

	var r0 []ovsdb.Operation
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]ovsdb.Operation, error)); ok {
		return rf(addresses)
	}
	if rf, ok := ret.Get(0).(func([]string) []ovsdb.Operation); ok {
		r0 = rf(addresses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ovsdb.Operation)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(addresses)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAddresses provides a mock function with given fields: addresses
func (_m *AddressSet) DeleteAddresses(addresses []string) error {
	ret := _m.Called(addresses)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAddresses")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(addresses)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAddressesReturnOps provides a mock function with given fields: addresses
func (_m *AddressSet) DeleteAddressesReturnOps(addresses []string) ([]ovsdb.Operation, error) {
	ret := _m.Called(addresses)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAddressesReturnOps")
	}

	var r0 []ovsdb.Operation
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]ovsdb.Operation, error)); ok {
		return rf(addresses)
	}
	if rf, ok := ret.Get(0).(func([]string) []ovsdb.Operation); ok {
		r0 = rf(addresses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ovsdb.Operation)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(addresses)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Destroy provides a mock function with given fields:
func (_m *AddressSet) Destroy() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Destroy")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetASHashNames provides a mock function with given fields:
func (_m *AddressSet) GetASHashNames() (string, string) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetASHashNames")
	}

	var r0 string
	var r1 string
	if rf, ok := ret.Get(0).(func() (string, string)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// GetAddresses provides a mock function with given fields:
func (_m *AddressSet) GetAddresses() ([]string, []string) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAddresses")
	}

	var r0 []string
	var r1 []string
	if rf, ok := ret.Get(0).(func() ([]string, []string)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func() []string); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	return r0, r1
}

// GetName provides a mock function with given fields:
func (_m *AddressSet) GetName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SetAddresses provides a mock function with given fields: addresses
func (_m *AddressSet) SetAddresses(addresses []string) error {
	ret := _m.Called(addresses)

	if len(ret) == 0 {
		panic("no return value specified for SetAddresses")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(addresses)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAddressSet creates a new instance of AddressSet. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAddressSet(t interface {
	mock.TestingT
	Cleanup(func())
}) *AddressSet {
	mock := &AddressSet{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
