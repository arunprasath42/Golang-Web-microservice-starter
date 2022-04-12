package mocks

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (UserService *MockUserService) CreateUsers(byteInput io.Reader) error {
	ret := UserService.Called(byteInput)

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r1
}

func (UserService *MockUserService) GetCommonFunc(byteInput io.Reader) error {
	ret := UserService.Called(byteInput)

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r1
}
