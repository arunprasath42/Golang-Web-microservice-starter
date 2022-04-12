package models

/*import (
	db "web-api/utils/database"
)

type Users struct {
	ID       int    `gorm:"column:id;type:int;primary_key;comment:Unique ID" json:"id"`
	Name     string `gorm:"column:name;type:varchar(255);comment:Name Info" json:"name"`
	Email    string `gorm:"column:email;type:varchar(255);comment:email Info" json:"email"`
	Location string `gorm:"column:location;type:varchar(255);comment:location Info" json:"location"`
}
func UserMigrate() {
	db.DB.AutoMigrate(&Users{})
}*/

type Users struct {
	Unique_id int `gorm:"column:unique_id; PRIMARY_KEY" json:"unique_id"`
	Name      string
	Email     string
	Location  string
}
