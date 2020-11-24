package entity

import (
	"time"
)

type Role struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:100;not null;" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	IsAdmin   bool      `gorm:"size:100;not null;" json:"is_admin" sql:"default:0"`
}

type Roles []Role
