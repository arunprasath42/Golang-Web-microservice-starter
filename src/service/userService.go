package service

import (
	"fmt"
	"web-api/src/models"
	"web-api/src/repository"
	"web-api/utils/constant"
	"web-api/utils/database"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type TestAPIUsers struct{}

/****API to insert users into database****/
func (c *TestAPIUsers) CreateUsers(insert models.Users) (string, error) {

	repo := repository.UserRepo{}
	result, err := repo.CreateUsers(insert)
	if err != nil {
		return "Please check the input", nil
	} else {
		fmt.Println("result", result)
		return "User Inserted Sucessfully", nil
	}
}

/***API to Read users from database****/
func (r *TestAPIUsers) FetchUsers(read models.Users) ([]models.Users, error) {
	var user []models.Users
	sqlDB := database.DB
	sqlDB.Debug().Where("name = ? ", read.Name).Find(&user)
	return user, nil
}

/****API to Update users from database****/
func (update *TestAPIUsers) EditFilter(edit models.Users) (string, error) {
	var val *gorm.DB
	sqlDB := database.DB
	log.Info().Msg(fmt.Sprintf("%v", sqlDB))

	if constant.Alphabets(edit.Name) && constant.Email(edit.Email) && constant.Alphabets(edit.Location) {
		val = sqlDB.Model(models.Users{}).Where("Unique_id IN (?)", edit.Unique_id).Updates(models.Users{Name: edit.Name, Email: edit.Email, Location: edit.Location})
		fmt.Println("value", val)
	} else {
		return "Check input for Edit Users", nil
	}
	if val.RowsAffected > 0 {
		return "User has been updated sucessfully", nil
	}
	return "User already exists", val.Error
}

/****API to Delete users from database****/
func (delete *TestAPIUsers) DeleteUser(del models.Users) (string, error) {
	var val *gorm.DB
	sqlDB := database.DB
	var User models.Users
	log.Info().Msg(fmt.Sprintf("%v", sqlDB))

	val = sqlDB.Debug().Where("Unique_id IN (?) ", del.Unique_id, del.Name, del.Email, del.Location).Delete(&User)
	fmt.Println("value", val)

	if val.RowsAffected > 0 {
		return "Status:Deleted Succesfully", nil
	}
	return "User doesn't exists or Already Deleted", val.Error
}
