// Code generated by mockery v2.3.0. DO NOT EDIT.

package resource

import mock "github.com/stretchr/testify/mock"

// MockSupplier is an autogenerated mock type for the Supplier type
type MockSupplier struct {
	mock.Mock
}

// Resources provides a mock function with given fields:
func (_m *MockSupplier) Resources() ([]Resource, error) {
	ret := _m.Called()

	var r0 []Resource
	if rf, ok := ret.Get(0).(func() []Resource); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Resource)
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
