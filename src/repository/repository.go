package repository

import (
	"web-api/utils/database"
)

var Repo MysqlRepository

type MySqlRepositoryRepo struct{}

func MySqlInit() {
	Repo = &MySqlRepositoryRepo{}
}

/***Inserting data to database***/
func (r *MySqlRepositoryRepo) Insert(req interface{}) error {
	if err := database.DB.Debug().Create(req).Error; err != nil {
		return err
	}
	return nil
}

/***Fetching data from database***/
func (r *MySqlRepositoryRepo) FindById(obj interface{}, id int) error {
	if err := database.DB.Debug().Where("unique_id = ? ", id).Find(obj).Error; err != nil {
		return err
	}
	return nil
}

/***For Updating tableName ***/
func (r *MySqlRepositoryRepo) Update(obj interface{}, id int, update interface{}) error {

	if err := database.DB.Debug().Where("Unique_id IN (?) ", id).First(obj).Updates(update).Error; err != nil {
		return err
	}
	return nil
}

func (r *MySqlRepositoryRepo) Delete(obj interface{}, id int) error {
	if err := database.DB.Debug().Where("Unique_id IN (?) ", id).First(obj).Delete(obj).Error; err != nil {
		return err
	}
	return nil
}
