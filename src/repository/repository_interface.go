package repository

import (
	"github.com/jinzhu/gorm"
)

//import "sondr-backend/src/models"

/***This interface is common for all the repository files****/
type MysqlRepository interface {
	Insert(req interface{}) error
	FindById(obj interface{}, id int) error
	Update(obj interface{}, id int, update interface{}) error
	Delete(obj interface{}, id int) error
	//CreateSubadmin(req *models.Admin) error // Admin Repository Interface
	GetAdmin(obj interface{}, email string) error //Read Admin data from database
	Find(obj interface{}, whereQuery string, value ...interface{}) error
	ListAllWithPagination(obj interface{}, selectQuery, tableName, joinsQuery string, pageno, pagesize int) (int, error)
	UpdateSubAdmin(obj interface{}, id int, update interface{}) *gorm.DB
}
