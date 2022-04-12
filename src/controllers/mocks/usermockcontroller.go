package mocks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"web-api/src/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/mock"
)

type MockUserController struct {
	mock.Mock
}

func ByteFunc(req interface{}) []byte {
	mockUserReq := req
	byteInput, err := json.Marshal(mockUserReq)
	if err != nil {
		fmt.Printf("err")
		log.Error().Msg(err.Error())
	}
	return byteInput
}

func UserMockFunc(byteInput []byte, user interface{}, funcname string) *mocks.MockUserService {
	res := user
	mockUserService := new(mocks.MockUserService)
	mockUserService.On(funcname, bytes.NewReader(byteInput)).Return(res, nil)
	return mockUserService
}

type FiltersMock struct {
	Name     string `default:"Arunprasath"`
	Email    string `default:"arun.ps@accubits.com"`
	Location string `default:"Coimbatore"`
}

//Mock function for Inserting users
func (mysqlController *MockUserController) SavedFilters(c *gin.Context) {
	mockUserReq := FiltersMock{}
	byteInput := ByteFunc(mockUserReq)
	mockAlertsTriggerService := UserMockFunc(byteInput, "User Inserted Sucessfully", "ReadUsers")
	err := mockAlertsTriggerService.CreateUsers(bytes.NewReader(byteInput))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"apiResponse": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"apiResponse": "User Inserted Sucessfully"})
	}
}
