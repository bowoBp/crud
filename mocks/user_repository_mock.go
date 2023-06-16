package mocks

import (
	"crud/entity"
	"github.com/stretchr/testify/mock"
)

// UserRepositoryMock is a mock implementation of the UserRepository interface
type UserRepositoryMock struct {
	mock.Mock
}

// GetUserById is a mock implementation of the GetUserById method
func (m *UserRepositoryMock) GetUserById(id uint) (entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(entity.User), args.Error(1)
}
