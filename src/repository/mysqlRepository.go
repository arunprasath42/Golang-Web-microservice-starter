package repository

import (
	"fmt"
	"sondr-backend/utils/database"
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
func (r *MySqlRepositoryRepo) FindById(obj interface{}, id uint) error {
	if err := database.DB.Debug().Where("id = ? ", id).Find(obj).Error; err != nil {
		return err
	}
	return nil
}

/***For Updating tableName ***/
func (r *MySqlRepositoryRepo) Update(obj interface{}, whereQuery string, id uint, update interface{}) error {

	if err := database.DB.Debug().Where(whereQuery, id).First(obj).Updates(update).Error; err != nil {
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
func (r *MySqlRepositoryRepo) ListAllWithPagination(obj interface{}, selectQuery, tableName, joinsQuery string, pageno, pagesize int) (int, error) {
	var count int
	var err error

	offset := (pageno - 1) * pagesize

	db := database.DB.Debug().Table(tableName)
	if selectQuery != "" {
		db = db.Select(selectQuery)
	}
	if joinsQuery != "" {
		db = db.Joins(joinsQuery)
	}
	if err = db.Count(&count).Limit(pagesize).Offset(offset).Find(obj).Error; err != nil {
		fmt.Println("error in db", err)
		return count, err
	}
	//	log.Println("event", event)
	return count, nil

}

func (r *MySqlRepositoryRepo) FindByIdWithJoins(obj interface{}, selectQuery, tableName, whereQuery, joinsQuery string, id uint) (int, error) {
	var count int
	var err error
	if err = database.DB.Debug().Table(tableName).Where(whereQuery, id).Select(selectQuery).Joins(joinsQuery).Count(&count).Scan(obj).Error; err != nil {
		fmt.Println("error in find", err)
		return count, err
	}
	fmt.Println("event", obj)
	return count, nil
}
