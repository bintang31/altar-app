package entity

import (
	"strings"
	"time"
)

//Transaction : Struct Entity Transaction
type Transaction struct {
	ID                   uint64    `gorm:"primary_key;auto_increment" json:"id"`
	TotalDrd             int64     `gorm:"size:100;not null;" json:"total_drd"`
	Denda                float64   `gorm:"size:100;not null;" json:"denda"`
	Pelanggan            string    `json:"pelanggan" gorm:"null; size:100"`
	Status               int       `gorm:"size:100;not null;" json:"status"`
	Total                float64   `gorm:"size:100;not null;" json:"total"`
	Pdam                 string    `json:"pdam" gorm:"not null; size:100"`
	Notes                string    `json:"notes" gorm:"null; size:100"`
	Jenis                string    `json:"jenis" gorm:"not null; size:100"`
	TotalAir             float64   `gorm:"size:100;not null;" json:"total_air"`
	TotalNonair          float64   `gorm:"size:100;not null;" json:"total_nonair"`
	LoketMessage         string    `json:"loket_message" gorm:"null; size:255"`
	LoketMessageCode     string    `json:"loket_message_code" gorm:"null; size:255"`
	LoketMessageStatus   string    `json:"loket_message_status" gorm:"null; size:255"`
	PeriodeBilling       string    `json:"periode_billing" gorm:"null; size:255"`
	PeriodeKolektif      string    `json:"periode_kolektif" gorm:"null; size:255"`
	TransactionsKolektif int       `json:"transactions_kolektif" gorm:"null"`
	CreatedBy            uint64    `gorm:"size:20;not null;" json:"created_by"`
	CreatedAt            time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt            time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
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
	Nosamb    string `gorm:"size:100;not null;" json:"nosamb"`
	Pin       int    `gorm:"size:100;not null;" json:"pin"`
	Pdam      string `gorm:"size:100;not null;" json:"pdam"`
	UserLoket string `gorm:"size:100;not null;" json:"user_loket"`
	Notes     string `gorm:"size:100;null;" json:"notes"`
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

// TransactionPelanggan : struct for transaction pelanggan
type TransactionPelanggan struct {
	ID              int     `json:"id" gorm:"primary_key"`
	Nosamb          string  `json:"nosamb"`
	Nama            string  `json:"nama"`
	Alamat          string  `json:"alamat"`
	Notes           string  `json:"notes"`
	Golongan        string  `json:"golongan" gorm:"not null; unique; size:255"`
	KodePdam        string  `json:"kode_pdam" gorm:"not null; unique; size:50"`
	Pdam            string  `json:"pdam" gorm:"not null; unique; size:50"`
	RayonName       string  `json:"rayon_name" gorm:"null"`
	StatusKolektif  string  `json:"status_kolektif" gorm:"null; size:100"`
	PeriodeBilling  string  `json:"periode_billing" gorm:"null; size:100"`
	StatusPelanggan string  `json:"status_pelanggan" gorm:"null; size:100"`
	StatusBilling   string  `json:"status_billing" gorm:"null; size:100"`
	TotalTagihan    float64 `json:"total_tagihan"`
	TglBayar        string  `json:"tgl_bayar"`
}

// RiwayatTransactionPelanggan : struct for transaction pelanggan
type RiwayatTransactionPelanggan struct {
	Periode         string  `json:"periode"`
	Bulan           string  `gorm:"size:100;not null;" json:"bulan"`
	Status          string  `gorm:"size:100;not null;" json:"status"`
	TotalPerPeriode float64 `gorm:"size:100;null;" json:"total_per_periode"`
	TotalPemakaian  float64 `gorm:"size:100;null;" json:"total_pemakaian"`
}
