package repository

import (
	"fmt"
	"web-api/src/models"
	"web-api/utils/database"

	"gorm.io/gorm"
)

/****ADD THE QUERY LOGIC HERE FOR EASE OF ACCESS IN THR SERVICE-FILE****/

type UserRepo struct{}

func (create *UserRepo) CreateUsers(c models.Users) (string, error) {
	var val *gorm.DB
	sqlDB := database.DB

	Users := models.Users{Name: c.Name, Email: c.Email, Location: c.Location}
	val = sqlDB.Debug().Select("Name", "Email", "Location").Create(&Users)
	fmt.Println("value", val)

	if val.RowsAffected > 0 {
		return "Status: User Saved", nil
	}
	return "null", val.Error
}

/***Need to be modified***/
func (read *UserRepo) ReadUsers(r models.Users) ([]models.Users, error) {
	var readUser []models.Users
	sqlDB := database.DB
	sqlDB.Debug().Where("name = ? ", r.Name).Find(&readUser)
	return readUser, nil
}
