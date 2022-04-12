package controllers

import (
	"fmt"
	"net/http"
	"web-api/src/models"
	service "web-api/src/service"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

/****API to Create data into the database***/
func CreateUsers(c *gin.Context) {
	var create = service.TestAPIUsers{}

	reqModel := models.Users{}
	if c.ShouldBind(&reqModel) == nil {
		log.Info().Msg("Binding successful for params-GetSavedUsers")
	}
	saved, err := create.CreateUsers(reqModel)

	fmt.Println("err", err)
	if err != nil {
		log.Error().Msgf("Error inserting data into the database: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"apiResponse": err})
	} else {
		log.Info().Msgf("User Inserted Succesfully in the database: %s", saved)
		c.JSON(http.StatusOK, gin.H{"apiResponse": saved})
	}
}

/****API to Read data from the database***/
func FetchUsers(c *gin.Context) {
	var UserService = service.TestAPIUsers{}
	reqModel := models.Users{}
	if c.ShouldBind(&reqModel) == nil {
		log.Info().Msg("binding successful for params-FetchUsers")
	}
	readUsers, err := UserService.FetchUsers(reqModel)
	fmt.Println("err", err)
	fmt.Println("Users", readUsers)

	if err != nil {
		log.Error().Msgf("Error while fetching data from the database: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"apiResponse": err})
	} else {
		log.Error().Msgf("Data found for ReadUsers")
		c.JSON(http.StatusOK, gin.H{"apiResponse": readUsers})
	}
}

/****API to Update data in the database***/
func EditUsers(c *gin.Context) {
	var EditFiltersService = service.TestAPIUsers{}
	fmt.Printf("in controller-EditFilters")
	reqDto := models.Users{}
	if c.ShouldBind(&reqDto) == nil {
		log.Info().Msg("binding successful for params--EditFilters")
	}
	saved, err := EditFiltersService.EditFilter(reqDto)
	fmt.Println("err", err)
	if err != nil {
		log.Error().Msgf("Error accessing data for EditFilters: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"apiResponse": err})
	} else {
		log.Error().Msgf("Updated successfully")
		c.JSON(http.StatusOK, gin.H{"apiResponse": saved})
	}
}

/****API to Delete data from the database***/
func DropFilters(c *gin.Context) {
	var AlertsTriggerService = service.TestAPIUsers{}
	fmt.Printf("in controller-DeleteFilters")
	reqDto := models.Users{}
	if c.ShouldBind(&reqDto) == nil {
		log.Info().Msg("binding successful for params-DropFilters")
	}
	drop, err := AlertsTriggerService.DeleteUser(reqDto)
	if err != nil {
		log.Error().Msgf("Error accessing data for DeleteFilters: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"apiResponse": err})
	} else {
		log.Error().Msgf("Data has been deleted for DeleteFilters")
		c.JSON(http.StatusOK, gin.H{"apiResponse": drop})
	}
}
