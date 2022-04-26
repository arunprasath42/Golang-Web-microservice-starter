package controllers

import (
	_ "github.com/go-sql-driver/mysql"
)

/****API to Create data into the database***/
// func CreateUsers(c *gin.Context) {

// 	reqModel := models.Users{}
// 	if err := c.ShouldBind(&reqModel); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
// 		return
// 	}
// 	if err := val.Validate(reqModel); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
// 		return
// 	}
// 	var service = service.TestAPIAdmin{}
// 	saved, err := service.CreateUsers(&reqModel)
// 	if err != nil {
// 		log.Error().Msgf("Error inserting data into the database: %s", err.Error())
// 		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
// 		return
// 	}
// 	log.Info().Msgf("User Inserted Succesfully in the database: %s", saved)
// 	c.JSON(http.StatusOK, response.SuccessResponse(saved))

// }

// /****API to Read data from the database***/
// func ReadUsers(c *gin.Context) {
// 	reqModel := &models.Users{}
// 	if err := c.ShouldBind(&reqModel); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
// 		return
// 	}
// 	if err := val.ValidateVariable(reqModel.Unique_id, "required", "unique_id"); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
// 		return
// 	}
// 	var service = service.TestAPIAdmin{}
// 	readUsers, err := service.FetchUsers(reqModel)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
// 		return
// 	}
// 	c.JSON(http.StatusOK, response.SuccessResponse(readUsers))
// }

// /****API to Update data in the database***/
// func UpdateUsers(c *gin.Context) {
// 	reqModel := &models.Users{}
// 	if err := c.ShouldBind(&reqModel); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
// 		return
// 	}
// 	if err := val.ValidateVariable(reqModel.Unique_id, "required", "unique_id"); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
// 		return
// 	}
// 	// if err := val.Validate(reqModel); err != nil {
// 	// 	c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
// 	// 	return
// 	// }
// 	var service = service.TestAPIAdmin{}
// 	saved, err := service.UpdateUsers(reqModel)
// 	if err != nil {
// 		log.Error().Msgf("Error updating data into the database: %s", err.Error())
// 		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
// 		return
// 	}
// 	log.Info().Msgf("User Updated Succesfully in the database: %s", saved)
// 	c.JSON(http.StatusOK, response.SuccessResponse(saved))

// }

// /****API to Delete data from the database***/
// func DeleteUsers(c *gin.Context) {
// 	reqModel := &models.Users{}
// 	if err := c.ShouldBind(&reqModel); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
// 		return
// 	}
// 	if err := val.ValidateVariable(reqModel.Unique_id, "required", "unique_id"); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
// 		return
// 	}
// 	var service = service.TestAPIAdmin{}
// 	deleted, err := service.DeleteUsers(reqModel)
// 	if err != nil {
// 		log.Error().Msgf("Error deleting data from the database: %s", err.Error())
// 		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
// 		return
// 	}
// 	log.Info().Msgf("Deleted successfull: %s", deleted)
// 	c.JSON(http.StatusOK, response.SuccessResponse(deleted))

// }
