package controllers

import (
	"net/http"
	"sondr-backend/src/models"
	service "sondr-backend/src/service"
	"sondr-backend/utils/constant"
	"sondr-backend/utils/response"
	val "sondr-backend/utils/validator"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

/****API to Create data into the database***/
func CreateSubadmin(c *gin.Context) {
	reqModel := models.Admin{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.TestAPIUsers{}
	saved, err := service.CreateSubadmin(&reqModel)
	if err != nil {
		log.Error().Msgf("Error inserting data into the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	log.Info().Msgf("User Inserted Succesfully in the database: %s", saved)
	c.JSON(http.StatusOK, response.SuccessResponse(saved))
}

/****API to Create data into the database***/
func GeneratePassword(c *gin.Context) {
	var service = service.TestAPIUsers{}
	generatePassword, err := service.GeneratePassword()
	if err != nil {
		log.Error().Msgf("Error in generating password: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	log.Info().Msgf("Password Generated: %s", generatePassword)
	c.JSON(http.StatusOK, response.SuccessResponse(generatePassword))
}
