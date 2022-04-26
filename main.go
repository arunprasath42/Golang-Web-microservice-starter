package main

import (
	"sondr-backend/route"
	"sondr-backend/src/repository"
	"sondr-backend/utils/database"
	"sondr-backend/utils/validator"

	config "sondr-backend/config"
	"sondr-backend/migration"
	logger "sondr-backend/utils/logging"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.LoadConfig()
	router := gin.Default()
	logger.NewLogger(viper.GetString("logging.level"))
	database.GetInstancemysql()
	migration.Migration()
	repository.MySqlInit()
	validator.Init()
	//router.Use(middleware.TracingMiddleware())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))
	route.SetupRoutes(router)

}
