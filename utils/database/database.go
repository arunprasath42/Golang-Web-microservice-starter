package database

import (
	"fmt"
	"web-api/utils/logging"

	"github.com/spf13/viper"

	"sync"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once
var DB *gorm.DB

//Conn connects to Database
func GetInstancemysql() (dba *gorm.DB) {
	once.Do(func() {
		user := viper.GetString("mysql.user")
		password := viper.GetString("mysql.password")
		host := viper.GetString("mysql.host")
		port := viper.GetString("mysql.port")
		dbname := viper.GetString("mysql.dbname")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
		db, err := gorm.Open(mysql.Open(dsn+"?parseTime=true"), &gorm.Config{})
		//close connection - cleanup and close
		dba = db
		if err != nil {
			log.Panic().Msgf("Error connecting to the database at %s:%s/%s", host, port, dbname)
			log.Info().Msgf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
		}

		sqlDB, err := dba.DB()
		if err != nil {
			log.Panic().Msgf("Error getting GORM DB definition")
		}
		sqlDB.SetMaxIdleConns(2)
		sqlDB.SetMaxOpenConns(10)
		//defer sqlDB.Close()
		logging.Logger.Info("Database connected successfully...")
		log.Info().Msgf("Successfully established connection to %s:%s/%s", host, port, dbname)

	})
	DB = dba
	return dba
}
