package repository

import "sondr-backend/src/models"

/***This interface is common for all the repository files****/
type MysqlRepository interface {
	Insert(req interface{}) error
	FindById(obj interface{}, id uint) error
	Update(obj interface{}, whereQuery string, id uint, update interface{}) error
	Delete(obj interface{}, id int) error
	CreateSubadmin(req *models.Admin) error // Admin Repository Interface
	//ListAllWithPagination()

	/***User Repository Interface***/

	/***Event Repository Interface***/
	ListAllWithPagination(obj interface{}, selectQuery, tableName, joinsQuery string, pageno, pagesize int) (int, error)
	FindByIdWithJoins(obj interface{}, selectQuery, tableName, whereQuery, joinsQuery string, id uint) (int, error)
}
