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

//TransactionsKolektif : Struct Entity TransactionKolektif
type TransactionsKolektif struct {
	ID           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	KodeKolektif string    `json:"kode_kolektif" gorm:"not null; size:100"`
	Pelanggan    int       `gorm:"size:100;not null;" json:"pelanggan"`
	TotalTagihan float64   `gorm:"size:100;not null;" json:"total_tagihan"`
	Pdam         string    `json:"pdam" gorm:"not null; size:100"`
	CreatedBy    uint64    `gorm:"size:20;not null;" json:"created_by"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	SisaTagihan  float64   `gorm:"size:100;not null;" json:"sisa_tagihan"`
}

//TransactionsKolektifs : Struct list TransactionKolektifs
type TransactionsKolektifs []TransactionsKolektif

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

// TransactionsBulkData : struct for transaction bulk
type TransactionsBulkData struct {
	SambunganRumah []RekairDetailBulk         `json:"sambungan_rumah"`
	Kolektif       []RekairDetailBulkKolektif `json:"kolektif"`
	Pin            int                        `json:"pin"`
}

// RekairDetailBulk : struct for transaction bulk
type RekairDetailBulk struct {
	Nosamb       string  `json:"nosamb"`
	TotalTagihan float64 `json:"total_tagihan"`
}

// RekairDetailBulkKolektif : struct for transaction bulk
type RekairDetailBulkKolektif struct {
	Kodekolektif      string                              `json:"kode_kolektif"`
	TotalPelanggan    int                                 `json:"total_pelanggan"`
	TotalTagihan      float64                             `json:"total_tagihan"`
	PelangganKolektif []RekairDetailBulkPelangganKolektif `json:"pelanggan_kolektif"`
}

// RekairDetailBulkPelangganKolektif : struct for transaction bulk
type RekairDetailBulkPelangganKolektif struct {
	Nosamb       string  `json:"nosamb"`
	TotalTagihan float64 `json:"total_tagihan"`
}
