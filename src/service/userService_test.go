package service

import (
	"bytes"
	"testing"
	"web-api/src/controllers/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

type ReadUserMock struct {
	Name     string `default:"Arunprasath"`
	Email    string `default:"arun.ps@accubits.com"`
	Location string `default:"india"`
}

func TestCreateUsers(t *testing.T) {
	mockUserReq := &ReadUserMock{}
	byteInput := mocks.ByteFunc(mockUserReq)
	result := mocks.UserMockFunc(byteInput, "User Inserted Sucessfully", "CreateUsers")
	err := result.CreateUsers(bytes.NewReader(byteInput))
	assert.NoError(t, err)
}
