package test

import (
	"fmt"
	"log"
	"testing"
	"web-api/src/models"
	"web-api/src/repository"
	"web-api/src/service"
	"web-api/utils/database"

	"github.com/spf13/viper"
)

func init() {
	fmt.Println("called")
	//config.LoadConfig()
	//logging.NewLogger(viper.GetString("log.filepath"), viper.GetString("log.level"))
	LoadConfig()
	database.GetInstancemysql()
	repository.MySqlInit()

}

func TestValidCreateUsers(t *testing.T) {

	args := models.Users{
		Name:     "abc",
		Email:    "abc@gmail.com",
		Location: "chennai",
	}
	s := service.TestAPIUsers{}
	_, err := s.CreateUsers(&args)
	if err != nil {
		t.Error("not valid")
	}

}

func TestInvalidCreateUsers(t *testing.T) {
	args := models.Users{
		Name:     "abc",
		Email:    "abc@gmail.com",
		Location: "chennai",
	}
	s := service.TestAPIUsers{}
	_, err := s.CreateUsers(&args)
	if err == nil {
		t.Error("not valid")
	}
}

func LoadConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("..")

	err := viper.ReadInConfig()
	if err != nil {
		log.Panic("Config not found...")
	}
}
