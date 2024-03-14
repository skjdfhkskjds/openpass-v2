// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

type DB_Expecter struct {
	mock *mock.Mock
}

func (_m *DB) EXPECT() *DB_Expecter {
	return &DB_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *DB) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type DB_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *DB_Expecter) Close() *DB_Close_Call {
	return &DB_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *DB_Close_Call) Run(run func()) *DB_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DB_Close_Call) Return(_a0 error) *DB_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_Close_Call) RunAndReturn(run func() error) *DB_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: key
func (_m *DB) Delete(key []byte) error {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type DB_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - key []byte
func (_e *DB_Expecter) Delete(key interface{}) *DB_Delete_Call {
	return &DB_Delete_Call{Call: _e.mock.On("Delete", key)}
}

func (_c *DB_Delete_Call) Run(run func(key []byte)) *DB_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *DB_Delete_Call) Return(_a0 error) *DB_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_Delete_Call) RunAndReturn(run func([]byte) error) *DB_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: key
func (_m *DB) Read(key []byte) ([]byte, error) {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) ([]byte, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type DB_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - key []byte
func (_e *DB_Expecter) Read(key interface{}) *DB_Read_Call {
	return &DB_Read_Call{Call: _e.mock.On("Read", key)}
}

func (_c *DB_Read_Call) Run(run func(key []byte)) *DB_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *DB_Read_Call) Return(_a0 []byte, _a1 error) *DB_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_Read_Call) RunAndReturn(run func([]byte) ([]byte, error)) *DB_Read_Call {
	_c.Call.Return(run)
	return _c
}

// Write provides a mock function with given fields: key, value
func (_m *DB) Write(key []byte, value []byte) error {
	ret := _m.Called(key, value)

	if len(ret) == 0 {
		panic("no return value specified for Write")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, []byte) error); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type DB_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - key []byte
//   - value []byte
func (_e *DB_Expecter) Write(key interface{}, value interface{}) *DB_Write_Call {
	return &DB_Write_Call{Call: _e.mock.On("Write", key, value)}
}

func (_c *DB_Write_Call) Run(run func(key []byte, value []byte)) *DB_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].([]byte))
	})
	return _c
}

func (_c *DB_Write_Call) Return(_a0 error) *DB_Write_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_Write_Call) RunAndReturn(run func([]byte, []byte) error) *DB_Write_Call {
	_c.Call.Return(run)
	return _c
}

// NewDB creates a new instance of DB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *DB {
	mock := &DB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
