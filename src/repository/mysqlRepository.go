package repository

import (
	"fmt"
	"sondr-backend/utils/database"
)

/***Fetching data from database***/
func (r *MySqlRepositoryRepo) FindById(obj interface{}, id int) error {
	if err := database.DB.Debug().Where("id = ? AND email = ? ", id).Find(obj).Error; err != nil {
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

func (r *MySqlRepositoryRepo) Find(obj interface{}, tableName string, selectQuery string, whereQuery string, value ...interface{}) error {
	db := database.DB.Debug().Table(tableName)
	if selectQuery != "" {
		db.Select(selectQuery)
	}
	if err := db.Where(whereQuery, value...).Find(obj).Error; err != nil {
		return err
	}
	return nil
}

/***For Pagination***/
func (r *MySqlRepositoryRepo) ListAllWithPagination(obj interface{}, selectQuery, tableName, joinsQuery string, pageno, pagesize int, whereQuery string, value ...interface{}) (int, error) {
	var count int
	var err error

	offset := (pageno - 1) * pagesize
	var res []interface{}

	db := database.DB.Debug().Table(tableName)
	if whereQuery != "" {
		for _, val := range value {
			if val != nil {
				fmt.Println("value", val)
				res = append(res, val)
			}
		}
		db = db.Where(whereQuery, res...)

	}
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
	return count, nil

}
