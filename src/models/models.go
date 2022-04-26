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

func UserMigrate() {
	db.DB.Debug().AutoMigrate(&Users{}, &Admin{})
}

type Admins struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"` //validatePasswordString:"password"`
}

type AdminEditprofile struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required" validate:"email"`
	ID    uint   `json:"uniqueID"`
}

/***TAble Name***/
type Admin struct {
	gorm.Model
	Name     string    `gorm:"column:name;type:varchar(255);unique" json:"name" binding:"required"`
	Email    string    `gorm:"column:email;type:varchar(255);unique" json:"email" binding:"required" validate:"email"`
	Password string    `gorm:"column:password;type:varchar(255);unique" json:"password" binding:"required" validatePasswordString:"password"`
	LastSeen time.Time `json:"last_seen"`
	Role     string    `gorm:"default:'Subadmin';type:enum('Admin','Subadmin')" json:"role,omitempty"`
}

type AdminResponse struct {
	ListAdmin []*ListAdmin `json:"listAdmin,omitempty"`
	Count     int          `json:"count,omitempty"`
}

/***For Listing admin table ***/
type ListAdmin struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Email       string    `json:"email,omitempty"`
	Password    string    `json:"password,omitempty"`
	CreatedDate time.Time `json:"createdDate,omitempty"`
}

type AdminsPagination struct {
	Name     string
	PageID   int
	PageSize int
}
