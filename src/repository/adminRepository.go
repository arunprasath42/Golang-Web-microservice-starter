package repository

import (
	"fmt"
	"sondr-backend/src/models"
	"sondr-backend/utils/database"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

/***Inserting data to database***/
func (r *MySqlRepositoryRepo) CreateSubadmin(req *models.Admin) error {
	Password, _ := HashPassword(req.Password)
	fmt.Println(Password)

	//HashPassword(req.Password)
	//if err := database.DB.Debug().Select("Name", "Email", "Password").Create(insert).Error; err != nil
	if err := database.DB.Debug().Create(req).Error; err != nil {
		return err
	}
	return nil
}
