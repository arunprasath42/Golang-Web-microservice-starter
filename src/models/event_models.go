package models

import (
	db "sondr-backend/utils/database"
	"time"

	"gorm.io/gorm"
)

type Events struct {
	gorm.Model
	HostUserID uint      `json:"hostUserId,omitempty"`
	EventName  string    `json:"eventName,omitempty"`
	Location   string    `json:"location,omitempty"`
	EventMode  string    `json:"eventMode,omitempty"`
	Password   string    `json:"password,omitempty"`
	Date       string    `json:"date,omitempty"`
	StartTime  time.Time `json:"startTime,omitempty"`
	EndTime    time.Time `json:"endTime,omitempty"`
	Reason     string    `json:"reason,omitempty"`
	Status     string    `json:"status,omitempty"`
}

type EventMetadatas struct {
	ID            int       `json:"id,omitempty"`
	EventId       uint      `json:"eventId,omitempty"`
	InvitedUserId uint      `gorm:"foreignKey:id" json:"invitedUserId,omitempty"`
	IsAttended    bool      `json:"isAttended,omitempty"`
	CheckIn       time.Time `json:"checkIn,omitempty"`
	CheckOut      time.Time `json:"checkOut,omitempty"`
}
type Response struct {
	EventResponse      []*ListEvent     `json:"listEvent,omitempty"`
	Count              int              `json:"count,omitempty"`
	Event              *Event           `json:"event,omitempty"`
	InvitedGuest       []*InvitedGuest  `json:"invitedGuest,omitempty"`
	AttendedGuest      []*AttendedGuest `json:"attendedGuest,omitempty"`
	InvitedGuestCount  int              `json:"invitedGuestCount,omitempty"`
	AttendedGuestCount int              `json:"attendedGuestCount,omitempty"`
	Message            string           `json:"message,omitempty"`
}
type AttendedGuest struct {
	InvitedUserId uint      `json:"invitedUserID,omitempty"`
	FirstName     string    `json:"firstName,omitempty"`
	LastName      string    `json:"lastName,omitempty"`
	CheckIn       time.Time `json:"checkIn,omitempty"`
	CheckOut      time.Time `json:"checkOut,omitempty"`
}
type InvitedGuest struct {
	InvitedUserId uint   `json:"invitedUserID,omitempty"`
	FirstName     string `json:"firstName,omitempty"`
	LastName      string `json:"lastName,omitempty"`
}
type ListEvent struct {
	ID        uint      `json:"id,omitempty"`
	EventName string    `json:"eventName,omitempty"`
	Location  string    `json:"location,omitempty"`
	Status    string    `json:"status,omitempty"`
	EventMode string    `json:"eventMode,omitempty"`
	Date      string    `json:"date,omitempty"`
	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
}
type Event struct {
	Id         uint      `json:"id,omitempty"`
	HostUserID uint      `json:"hostUserId,omitempty"`
	EventName  string    `json:"eventName,omitempty"`
	Location   string    `json:"location,omitempty"`
	EventMode  string    `json:"eventMode,omitempty"`
	Password   string    `json:"password,omitempty"`
	Date       string    `json:"date,omitempty"`
	StartTime  time.Time `json:"startTime,omitempty"`
	EndTime    time.Time `json:"endTime,omitempty"`
	Status     string    `json:"status,omitempty"`
	FirstName  string    `json:"firstName,omitempty"`
	LastName   string    `json:"lastName,omitempty"`
}

func EventMigrate() {
	db.DB.AutoMigrate(&Events{}).AddForeignKey("host_user_id", "users(id)", "CASCADE", "CASCADE")
	db.DB.AutoMigrate(&EventMetadatas{}).AddForeignKey("event_id", "events(id)", "CASCADE", "CASCADE").AddForeignKey("invited_user_id", "users(id)", "CASCADE", "CASCADE")
}
