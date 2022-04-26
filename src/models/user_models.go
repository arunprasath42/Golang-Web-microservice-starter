package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	FirstName      string    `gorm:"type:varchar(255);" json:"first_name,omitempty" valid:"required,length(3|255)"`
	LastName       string    `gorm:"type:varchar(255);" json:"last_name,omitempty" valid:"required,length(2|255)"`
	Email          string    `gorm:"type:varchar(255);" json:"email,omitempty" valid:"required,length(5|255)"`
	Phone_no       string    `gorm:"type:varchar(25);" json:"phone_no,omitempty" valid:"required,length(10|25)"`
	Gender         string    `json:"gender,omitempty"`
	DOB            string    `json:"dob,omitempty"`
	ProfilePhoto   string    `json:"profile_photo,omitempty"`
	Occupation     string    `json:"occupation,omitempty"`
	Height         float64   `json:"height,omitempty"`
	Country        string    `gorm:"type:varchar(50);" json:"country,omitempty" valid:"required,length(10|50)"`
	City           string    `gorm:"type:varchar(50);" json:"city,omitempty" valid:"required,length(10|50)"`
	UploadedPhotos string    `json:"uploaded_photos,omitempty"`
	LastSeen       time.Time `json:"last_seen,omitempty"`
	Active         bool      `gorm:"type:bool;default:true" json:"active,omitempty"`
	Visible        bool      `gorm:"type:bool;default:true" json:"visible,omitempty"`
	LastVisited    time.Time `json:"last_visited,omitempty"`
	Blocked        bool      `gorm:"type:bool;default:false" json:"blocked,omitempty"`
	BlockerID      int       `json:"blocker_id,omitempty"`
	BlockerType    string    `json:"blocker_type,omitempty"`
}
