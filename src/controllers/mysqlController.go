package controllers

import (
	"fmt"
	"net/http"
	"web-api/src/controllers/dto"
	service "web-api/src/service"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

func MessageSlack(c *gin.Context) {
	fmt.Printf("Runnning Inside controller")
	reqDto := dto.TestSlackMessage{}
	if c.ShouldBind(&reqDto) == nil {
		log.Info().Msg("binding successful for params")
	}
	testslackService := service.TestAPI{}
	message, err := testslackService.TestMessageForSlack(reqDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"apiResponse": err})
	} else {
		fmt.Println(message)
		c.JSON(http.StatusOK, gin.H{"apiResponse": message})
	}
}
