package route

import (
	"web-api/src/controllers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupRoutes(router *gin.Engine) {

	/***BASEPATH OF AN API. NOTE:THIS SHOULDN'T BE CHANGED***/
	api := router.Group("/MocktestAPI")

	/***ADD THE ROUTES HERE***/
	api.POST("/saveUsers", controllers.CreateUsers)
	api.GET("/fetchUsers", controllers.FetchUsers)
	api.PUT("/UpdateUsers", controllers.EditUsers)
	api.POST("/deleteUsers", controllers.DropFilters)

	router.Run(viper.GetString("server.port"))
}
