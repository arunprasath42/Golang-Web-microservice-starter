package repository

import (
	//"sondr-backend/src/models"

	"sondr-backend/utils/database"

	"github.com/jinzhu/gorm"
)

/***Inserting Sub admins to database***/
// func (r *MySqlRepositoryRepo) (req *models.Admin) error {
// 	if err := database.DB.Debug().Create(req).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
var Repo MysqlRepository

type MySqlRepositoryRepo struct{}

func MySqlInit() {
	Repo = &MySqlRepositoryRepo{}
}

/**********************Creating SUB ADMINS****************************/
func (r *MySqlRepositoryRepo) Insert(req interface{}) error {
	if err := database.DB.Debug().Create(req).Error; err != nil {
		return err
	}
	return nil
}

/***********************Reading admin data from database**************************/
func (r *MySqlRepositoryRepo) GetAdmin(obj interface{}, email string) error {
	if err := database.DB.Debug().Where("email = ? ", email).Find(obj).Error; err != nil {
		return err
	}
	return nil
}

/***********************Updating sub-admin******************/
/*func (r *MySqlRepositoryRepo) UpdateSubAdmin(obj interface{}, id int, update interface{}) *gorm.DB {

	if err := database.DB.Debug().Where("id IN (?) ", id).First(obj).Updates(update); err != nil {
		return err
	}
	return nil
}*/
func (r *MySqlRepositoryRepo) UpdateSubAdmin(obj interface{}, id int, update interface{}) *gorm.DB {

	if err := database.DB.Debug().Model(obj).Where("id IN (?)", id).Updates(update); err != nil {
		return err
	}
	return nil
}
