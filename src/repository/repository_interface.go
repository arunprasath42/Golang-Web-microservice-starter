package repository

import "sondr-backend/src/models"

/***This interface is common for all the repository files****/
type MysqlRepository interface {
	Insert(req interface{}) error
	FindById(obj interface{}, id int) error
	Update(obj interface{}, id int, update interface{}) error
	Delete(obj interface{}, id int) error
	CreateSubadmin(req *models.Admin) error // Admin Repository Interface

	/***User Repository Interface***/

	/***Event Repository Interface***/
}
