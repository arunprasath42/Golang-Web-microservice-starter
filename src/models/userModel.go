package models

import (
	"time"
	db "web-api/utils/database"

	"gorm.io/gorm"
)

type Users struct {
	Unique_id int    `gorm:"column:unique_id; PRIMARY_KEY" json:"unique_id"`
	Name      string `gorm:"column:name;type:varchar(255);unique" json:"name"`
	Email     string `gorm:"column:email;type:varchar(255);unique" json:"email" validate:"email"`
	Location  string `gorm:"column:location;type:varchar(255);" json:"location"`
}

type Admin struct {
	gorm.Model
	Name      string    `gorm:"type:varchar(255);" json:"name,omitempty" validate:"required,length(3|255)"`
	Email     string    `gorm:"type:varchar(255);" json:"email,omitempty" validate:"required,length(5|255)"`
	Password  string    `gorm:"type:varchar(25);" json:"password,omitempty" validate:"required,length(8|25)"`
	LastSeen  time.Time `json:"last_seen,omitempty"`
	Role      string    `gorm:"default:'Admin';type:enum('Admin','Subadmin')" json:"role,omitempty" validate:"required,in(Admin|Subadmin)"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

func UserMigrate() {
	db.DB.Debug().AutoMigrate(&Users{}, &Admin{})
}
