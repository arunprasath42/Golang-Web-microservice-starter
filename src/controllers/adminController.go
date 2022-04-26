package controllers

import (
	//"fmt"

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

/*******************************CREATING SUB-ADMINS**********************************/

func CreateSubadmin(c *gin.Context) {

	// role := c.GetString("role")
	// if role != "Subadmin" {
	// 	c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, errors.New("Permission Denied")))
	// 	return
	// }

	reqModel := models.Admins{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.TestAPIAdmin{}
	saved, err := service.CreateSubadmin(&reqModel)
	if err != nil {
		log.Error().Msgf("Error inserting data into the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	log.Info().Msgf("User Inserted Succesfully in the database: %s", saved)
	c.JSON(http.StatusOK, response.SuccessResponse(saved))
}

/******************************GENARATE PASSWORD MANUALLY***************************/
func GeneratePassword(c *gin.Context) {
	var service = service.TestAPIAdmin{}
	generatePassword, err := service.GeneratePassword()
	if err != nil {
		log.Error().Msgf("Error in generating password: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	log.Info().Msgf("Password Generated: %s", generatePassword)
	c.JSON(http.StatusOK, response.SuccessResponse(generatePassword))
}

/*********************************LOGIN*********************************************/
func Login(c *gin.Context) {
	reqModel := models.Request{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.TestAPIAdmin{}
	saved, err := service.Login(&reqModel)
	if err != nil {
		log.Error().Msgf("Unable to Login: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	log.Info().Msgf("Login sucessfull: %s", saved)
	c.JSON(http.StatusOK, response.SuccessResponse(saved))

}

/*********************************LISTING SUBADMINS*********************************************/
func ListSubAdmins(c *gin.Context) {
	reqModel := &models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.TestAPIAdmin{}
	Subadmin, err := service.ListSubAdmins(reqModel.PageNo, reqModel.PageSize, reqModel.SearchFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(Subadmin))
}

/*func ListSubAdmins(c *gin.Context) {
	var service = service.TestAPIAdmin{}
	reqDto := models.AdminsPagination{}
	if c.ShouldBind(&reqDto) == nil {
		log.Info().Msg("binding successful for params-GetAlertsFilter")
	}
	Admins, totalrecords, err := service.ListSubAdmins(reqDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"SubAdminsList": err})
	} else {
		log.Error().Msgf("Requested data found for fetchFilters")
		c.JSON(http.StatusOK, gin.H{"SubAdminsList": Admins, "count": totalrecords})
	}
}*/

/*********************************READ SUBADMIN DETAILS*********************************************/
/*func ReadSubAdmin(c *gin.Context) {


	reqModel := &models.Admins{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.ValidateVariable(reqModel.ID, "required", "unique_id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.TestAPIAdmin{}
	readAdmins, err := service.ReadSubAdmin(reqModel)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(readAdmins))
}*/

/*********************************READ SUBADMIN DETAILS*********************************************/
func ReadSubAdmin(c *gin.Context) {

	reqModel := &models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.ValidateVariable(reqModel.ID, "required", "id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.TestAPIAdmin{}
	readAdmins, err := service.ReadSubAdmin(reqModel.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(readAdmins))
}

/*********************************UPDATE SUBADMIN DETAILS*********************************************/
// func UpdateSubAdmin(c *gin.Context) {
// 	reqModel := &models.Admins{}
// 	if err := c.ShouldBind(&reqModel); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
// 		return
// 	}
// 	if err := val.ValidateVariable(reqModel.ID, "required", "id"); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
// 		return
// 	}
// 	var updateSubadmin = service.TestAPIAdmin{}
// 	updated, err := updateSubadmin.UpdateSubAdmin(reqModel)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
// 		return
// 	}
// 	c.JSON(http.StatusOK, response.SuccessResponse(updated))
// }

// /******************************VERIFY PASSWORD*********************************************/
// func VerifyPassword(c *gin.Context) {
// 	reqModel := models.Admins{}
// 	if err := c.ShouldBindQuery(&reqModel); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
// 		return
// 	}
// 	if err := val.Validate(reqModel); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
// 		return
// 	}
// 	var service = service.TestAPIAdmin{}
// 	saved, err := service.VerifyPassword(&reqModel)
// 	if err != nil {
// 		log.Error().Msgf("Unable to Verify: %s", err.Error())
// 		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
// 		return
// 	}
// 	log.Info().Msgf("Verify sucessfull: %s", saved)
// 	c.JSON(http.StatusOK, response.SuccessResponse(saved))
// }

// /*********************************CHANGE PASSWORD*********************************************/

// func ChangePassword(c *gin.Context) {
// 	reqModel := &models.Admins{}

// 	if err := c.ShouldBind(&reqModel); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
// 		return
// 	}
// 	if err := val.Validate(reqModel); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
// 		return
// 	}
// 	var service = service.TestAPIAdmin{}
// 	changePassword, err := service.ChangePassword(reqModel)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
// 		return
// 	}
// 	c.JSON(http.StatusOK, response.SuccessResponse(changePassword))
// }
