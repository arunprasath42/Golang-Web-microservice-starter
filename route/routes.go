package route

import (
	controllers2 "web-api/src/controllers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupRoutes(router *gin.Engine) {

	api := router.Group(viper.GetString("server.basepath"))

	//api.POST(viper.GetString("server.jiraAuthAPI"), controllers2.JiraAuthenticate)
	api.POST(viper.GetString("server.slacktestmessageAPI"), controllers2.MessageSlack)

	router.Run(viper.GetString("server.port"))
}
