package service

import (
	"sondr-backend/src/models"
	"sondr-backend/src/repository"
)

type TestAPIAdmin struct{}

/****API to insert users into database****/
func (c *TestAPIAdmin) CreateUsers(insert *models.Users) (string, error) {

	if err := repository.Repo.Insert(insert); err != nil {
		return "Unable to insert user", err
	}

	return "User inserted sucessfully", nil
}

/***API to Read users from database****/
func (c *TestAPIAdmin) FetchUsers(read *models.Users) (*models.Users, error) {
	var user models.Users
	if err := repository.Repo.FindById(&user, read.Unique_id); err != nil {
		return nil, err
	}
	return &user, nil
}

/****API to Update users from database****/
func (c *TestAPIAdmin) UpdateUsers(edit *models.Users) (string, error) {

	if err := repository.Repo.Update(&models.Users{}, edit.Unique_id, edit); err != nil {
		return "Unable to update user", err
	}

	return "Update Sucessfull", nil
}

/****API to Delete users from database****/
func (c *TestAPIAdmin) DeleteUsers(del *models.Users) (string, error) {
	if err := repository.Repo.Delete(del, del.Unique_id); err != nil {
		return "Unable to delete user", err
	}
	return "Delete Sucessfull", nil
}
