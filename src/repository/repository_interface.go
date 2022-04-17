package repository

type MysqlRepository interface {
	Insert(req interface{}) error
	FindById(obj interface{}, id int) error
	Update(obj interface{}, id int, update interface{}) error
	Delete(obj interface{}, id int) error
}
