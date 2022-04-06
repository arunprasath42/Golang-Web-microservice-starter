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
	api.POST("/sendTestMessageToSlack", controllers.MessageSlack)

	router.Run(viper.GetString("server.port"))
}
