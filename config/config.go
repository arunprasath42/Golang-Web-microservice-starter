package config

import (
	"fmt"

	//applogger "web-api/utils/logging"

	//"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if readErr, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Error()
			fmt.Println(readErr)
		} else {
			log.Error()
		}
	}
}
