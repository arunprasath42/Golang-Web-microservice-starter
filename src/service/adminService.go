package service

import (
	"fmt"
	"log"
	"sondr-backend/src/models"
	"sondr-backend/src/repository"

	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

/****API to insert users into database****/
func (c *TestAPIUsers) CreateSubadmin(insert *models.Admin) (string, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(insert.Password), 4)
	insert.Password = string(password)

	if err := repository.Repo.CreateSubadmin(insert); err != nil {
		return "Unable to create Sub-admin", err
	}

	return "Sub Admin created successfully", nil
}

/****API to Generate password****/
func (c *TestAPIUsers) GeneratePassword() (string, error) {
	res, err := password.Generate(12, 1, 1, false, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	return res, nil
}
