package repository

import (
	"crud/entity"
	"github.com/stretchr/testify/mock"
)

type MockUserInterface struct {
	mock.Mock
}

func (_m *MockUserInterface) CreateUser(user *entity.User) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (_m *MockUserInterface) UpdateUser(user *entity.User, id uint) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (_m *MockUserInterface) DeleteUser(email string) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (_m *MockUserInterface) GetUserById(id uint) (entity.User, error) {
	ret := _m.Called(id)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(int) entity.User); ok {
		r0 = rf(int(id))
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(int(id))
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockUCRProductInterface interface {
	mock.TestingT
	Cleanup(func())
}

func NewMockUserInterface(t mockConstructorTestingTNewMockUCRProductInterface) *MockUserInterface {
	mock := &MockUserInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
