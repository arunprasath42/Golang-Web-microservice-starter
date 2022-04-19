package models

import (
	db "sondr-backend/utils/database"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Unique_id int    `gorm:"column:unique_id; PRIMARY_KEY" json:"unique_id"`
	Name      string `gorm:"column:name;type:varchar(255);unique" json:"name"`
	Email     string `gorm:"column:email;type:varchar(255);unique" json:"email" validate:"email"`
	Location  string `gorm:"column:location;type:varchar(255);" json:"location"`
}

// type Admin struct {
// 	gorm.Model
// 	Name     string    `gorm:"type:varchar(255);unique" json:"name,omitempty" validate:"required,length(3|255)"`
// 	Email    string    `gorm:"type:varchar(255);unique" json:"email,omitempty" validate:"required,length(5|255)"`
// 	Password string    `gorm:"type:varchar(25);unique" json:"password,omitempty" validate:"required,length(8|25)"`
// 	LastSeen time.Time `json:"last_seen,omitempty"`
// 	Role     string    `gorm:"default:'Admin';type:enum('Admin','Subadmin')" json:"role,omitempty" validate:"required,in(Admin|Subadmin)"`
// }

type Admin struct {
	gorm.Model
	Name     string    `gorm:"column:name;type:varchar(255);unique" json:"name" binding:"required"`
	Email    string    `gorm:"column:email;type:varchar(255);unique" json:"email" binding:"required" validate:"email"`
	Password string    `gorm:"column:password;type:varchar(255);unique" json:"password" binding:"required" validatePasswordString:"password"`
	LastSeen time.Time `json:"last_seen"`
	Role     string    `gorm:"default:'Admin';type:enum('Admin','Subadmin')" json:"role,omitempty"`
}

func UserMigrate() {
	db.DB.Debug().AutoMigrate(&Users{}, &Admin{})
}
