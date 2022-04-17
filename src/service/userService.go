package service

import (
	"web-api/src/models"
	"web-api/src/repository"
)

type TestAPIUsers struct{}

/****API to insert users into database****/
func (c *TestAPIUsers) CreateUsers(insert *models.Users) (string, error) {

	if err := repository.Repo.Insert(insert); err != nil {
		return "Unable to insert user", err
	}

	return "User inserted sucessfully", nil
}

/***API to Read users from database****/
func (c *TestAPIUsers) FetchUsers(read *models.Users) (*models.Users, error) {
	var user models.Users
	if err := repository.Repo.FindById(&user, read.Unique_id); err != nil {
		return nil, err
	}
	return &user, nil
}

/****API to Update users from database****/
func (c *TestAPIUsers) UpdateUsers(edit *models.Users) (string, error) {

	if err := repository.Repo.Update(&models.Users{}, edit.Unique_id, edit); err != nil {
		return "Unable to update user", err
	}

	return "Update Sucessfull", nil
}

/****API to Delete users from database****/
func (c *TestAPIUsers) DeleteUsers(del *models.Users) (string, error) {
	if err := repository.Repo.Delete(del, del.Unique_id); err != nil {
		return "Unable to delete user", err
	}
	return "Delete Sucessfull", nil
}
