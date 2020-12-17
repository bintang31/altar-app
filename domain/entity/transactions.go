package entity

import (
	"strings"
	"time"
)

//Transaction : Struct Entity Transaction
type Transaction struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	TotalDrd  float64   `gorm:"size:100;not null;" json:"total_drd"`
	Denda     float64   `gorm:"size:100;not null;" json:"denda"`
	Pelanggan string    `json:"pelanggan" gorm:"null; size:100"`
	Status    int       `gorm:"size:100;not null;" json:"status"`
	Total     float64   `gorm:"size:100;not null;" json:"total"`
	Pdam      string    `json:"pdam" gorm:"not null; size:100"`
	Jenis     string    `json:"jenis" gorm:"not null; size:100"`
	CreatedBy uint64    `gorm:"size:20;not null;" json:"created_by"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

//Transactions : Struct list Transaction
type Transactions []Transaction

//Bayar : Struct Bayar
type Bayar struct {
	Nosamb string `gorm:"size:100;not null;" json:"nosamb"`
	Pin    int    `gorm:"size:100;not null;" json:"pin"`
	Pdam   string `gorm:"size:100;not null;" json:"pdam"`
}

//Validate : user validation by Action
func (b *Bayar) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)

	switch strings.ToLower(action) {
	case "validate_pin":
		if b.Pin == 0 {
			errorMessages["email_required"] = "pin required"
		}

	default:
		if b.Pin == 0 {
			errorMessages["firstname_required"] = "pin  is required"
		}

	}
	return errorMessages
}
