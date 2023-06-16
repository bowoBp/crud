// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	user "crud/modules/user"

	mock "github.com/stretchr/testify/mock"
)

// ControllerUser is an autogenerated mock type for the ControllerUser type
type ControllerUser struct {
	mock.Mock
}

type ControllerUser_Expecter struct {
	mock *mock.Mock
}

func (_m *ControllerUser) EXPECT() *ControllerUser_Expecter {
	return &ControllerUser_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: req
func (_m *ControllerUser) CreateUser(req user.UserParam) (interface{}, error) {
	ret := _m.Called(req)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(user.UserParam) (interface{}, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(user.UserParam) interface{}); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(user.UserParam) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerUser_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type ControllerUser_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - req user.UserParam
func (_e *ControllerUser_Expecter) CreateUser(req interface{}) *ControllerUser_CreateUser_Call {
	return &ControllerUser_CreateUser_Call{Call: _e.mock.On("CreateUser", req)}
}

func (_c *ControllerUser_CreateUser_Call) Run(run func(req user.UserParam)) *ControllerUser_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(user.UserParam))
	})
	return _c
}

func (_c *ControllerUser_CreateUser_Call) Return(_a0 interface{}, _a1 error) *ControllerUser_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ControllerUser_CreateUser_Call) RunAndReturn(run func(user.UserParam) (interface{}, error)) *ControllerUser_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUser provides a mock function with given fields: email
func (_m *ControllerUser) DeleteUser(email string) (interface{}, error) {
	ret := _m.Called(email)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (interface{}, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerUser_DeleteUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUser'
type ControllerUser_DeleteUser_Call struct {
	*mock.Call
}

// DeleteUser is a helper method to define mock.On call
//   - email string
func (_e *ControllerUser_Expecter) DeleteUser(email interface{}) *ControllerUser_DeleteUser_Call {
	return &ControllerUser_DeleteUser_Call{Call: _e.mock.On("DeleteUser", email)}
}

func (_c *ControllerUser_DeleteUser_Call) Run(run func(email string)) *ControllerUser_DeleteUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ControllerUser_DeleteUser_Call) Return(_a0 interface{}, _a1 error) *ControllerUser_DeleteUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ControllerUser_DeleteUser_Call) RunAndReturn(run func(string) (interface{}, error)) *ControllerUser_DeleteUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserById provides a mock function with given fields: id
func (_m *ControllerUser) GetUserById(id uint) (user.FindUser, error) {
	ret := _m.Called(id)

	var r0 user.FindUser
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (user.FindUser, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) user.FindUser); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(user.FindUser)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerUser_GetUserById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserById'
type ControllerUser_GetUserById_Call struct {
	*mock.Call
}

// GetUserById is a helper method to define mock.On call
//   - id uint
func (_e *ControllerUser_Expecter) GetUserById(id interface{}) *ControllerUser_GetUserById_Call {
	return &ControllerUser_GetUserById_Call{Call: _e.mock.On("GetUserById", id)}
}

func (_c *ControllerUser_GetUserById_Call) Run(run func(id uint)) *ControllerUser_GetUserById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *ControllerUser_GetUserById_Call) Return(_a0 user.FindUser, _a1 error) *ControllerUser_GetUserById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ControllerUser_GetUserById_Call) RunAndReturn(run func(uint) (user.FindUser, error)) *ControllerUser_GetUserById_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: req, id
func (_m *ControllerUser) UpdateUser(req user.UserParam, id uint) (interface{}, error) {
	ret := _m.Called(req, id)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(user.UserParam, uint) (interface{}, error)); ok {
		return rf(req, id)
	}
	if rf, ok := ret.Get(0).(func(user.UserParam, uint) interface{}); ok {
		r0 = rf(req, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(user.UserParam, uint) error); ok {
		r1 = rf(req, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerUser_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type ControllerUser_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - req user.UserParam
//   - id uint
func (_e *ControllerUser_Expecter) UpdateUser(req interface{}, id interface{}) *ControllerUser_UpdateUser_Call {
	return &ControllerUser_UpdateUser_Call{Call: _e.mock.On("UpdateUser", req, id)}
}

func (_c *ControllerUser_UpdateUser_Call) Run(run func(req user.UserParam, id uint)) *ControllerUser_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(user.UserParam), args[1].(uint))
	})
	return _c
}

func (_c *ControllerUser_UpdateUser_Call) Return(_a0 interface{}, _a1 error) *ControllerUser_UpdateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ControllerUser_UpdateUser_Call) RunAndReturn(run func(user.UserParam, uint) (interface{}, error)) *ControllerUser_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewControllerUser creates a new instance of ControllerUser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewControllerUser(t interface {
	mock.TestingT
	Cleanup(func())
}) *ControllerUser {
	mock := &ControllerUser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
