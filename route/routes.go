package route

import (
	controllers "sondr-backend/src/controllers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupRoutes(router *gin.Engine) {

	/***BASEPATH OF AN API. NOTE:THIS SHOULDN'T BE CHANGED***/
	api := router.Group("/MocktestAPI")

	/***ADD THE ROUTES HERE***/
	api.POST("/saveUsers", controllers.CreateUsers)
	api.GET("/readUsers", controllers.ReadUsers)
	api.PUT("/updateUsers", controllers.UpdateUsers)
	api.DELETE("/deleteUsers", controllers.DeleteUsers)

	/***ADD THE ADMIN ROUTES***/
	api.POST("/createSubadmin", controllers.CreateSubadmin)
	api.POST("/generatePassword", controllers.GeneratePassword)
	api.POST("/login", controllers.Login)
	api.GET("/readSubadmin", controllers.ListSubAdmins)
	api.GET("/fetchAdmindetails", controllers.ReadSubAdmin)
	api.PUT("/updateSubadmin", controllers.UpdateSubAdmin)
	api.GET("/verifyPassword", controllers.VerifyPassword)
	api.PUT("/changePassword", controllers.ChangePassword)

	router.Run(viper.GetString("server.port"))
}
