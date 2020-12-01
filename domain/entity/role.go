package entity

import (
	"time"
)

//Role : Struct Entity Role
type Role struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:100;not null;" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	IsAdmin   bool      `gorm:"size:100;not null;" json:"is_admin" sql:"default:0"`
}

//Roles : List Struct Roles
type Roles []Role
