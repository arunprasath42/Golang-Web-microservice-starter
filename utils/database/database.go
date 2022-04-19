package database

import (
	"fmt"
	"sondr-backend/utils/logging"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var once sync.Once
var DB *gorm.DB

//Conn connects to Database
func GetInstancemysql() (dba *gorm.DB) {
	once.Do(func() {
		var mysqlHost = fmt.Sprint(viper.GetString("mysql.user"), ":", viper.GetString("mysql.password"), "@(", viper.GetString("mysql.host"), ")/", viper.GetString("mysql.dbname"), "?parseTime=true")
		db, err := gorm.Open("mysql", mysqlHost)
		if err != nil {
			logging.Logger.WithError(err).WithField("err", err).Errorf("Database not connected")
			panic(err)
		}
		DB = db
		DB.DB().SetMaxIdleConns(10)
		DB.DB().SetMaxOpenConns(40)
		logging.Logger.Info("Database connected successfully...")

	})

	return DB
}
