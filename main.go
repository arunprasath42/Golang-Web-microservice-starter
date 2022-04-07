package main

import (
	"flag"
	"fmt"
	"web-api/route"
	"web-api/utils/database"
	applogger "web-api/utils/logging"
	"web-api/utils/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

var configFilePath *string
var log *zerolog.Logger = applogger.GetInstance()

func main() {
	loadConfig()
	router := gin.Default()
	setupLogger(router)
	database.GetInstancemysql()
	router.Use(middleware.TracingMiddleware())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))
	route.SetupRoutes(router)

}

/***loadConfig - Load the config parameters from the conf/app.yaml***/
func loadConfig() {
	configFilePath = flag.String("config-path", "conf/", "conf/")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(*configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		if readErr, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Error()
			fmt.Println(readErr)
		} else {
			log.Error()
		}
	}
}

// setupLogger - Configure logging for the server
func setupLogger(r *gin.Engine) {
	// Configure logger
	zerologger := applogger.GetInstance()
	r.Use(logger.SetLogger(logger.Config{
		Logger: zerologger,
		UTC:    true,
	}))
}
