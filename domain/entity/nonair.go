package entity

import (
	"time"
)

//Nonair : Struct Entity Nonair
type Nonair struct {
	Nomor            string    `gorm:"size:100;not null;" json:"nomor"`
	Periode          string    `gorm:"size:100;not null;" json:"periode"`
	Urutan           string    `gorm:"size:255;not null;" json:"urutan,omitempty"`
	Jenis            string    `gorm:"size:100;not null;" json:"jenis,omitempty"`
	JenisTagihan     string    `gorm:"size:100;not null;" json:"jenis_tagihan"`
	Pdam             string    `gorm:"size:100;not null;" json:"pdam"`
	Angsur           string    `gorm:"size:100;null;" json:"angsur"`
	Bulan            string    `gorm:"size:100;null;" json:"bulan"`
	NoAngsuran       string    `gorm:"size:100;null;" json:"no_angsuran"`
	Termin           int       `json:"termin" gorm:"null"`
	KetJenis         string    `json:"ket_jenis" gorm:"null; size:20"`
	DibebankanKepada string    `json:"dibebankan_kepada" gorm:"null; size:20"`
	Flag             string    `gorm:"size:100;null;" json:"flag"`
	Lunas            string    `gorm:"size:100;null;" json:"lunas"`
	Total            float64   `gorm:"size:100;null;" json:"total"`
	Administrasi     float64   `gorm:"size:100;null;" json:"administrasi"`
	Biayapasang      float64   `gorm:"size:100;null;" json:"biayapasang"`
	DendaTunggakan   float64   `gorm:"size:100;null;" json:"denda_tunggakan"`
	Lainnya          float64   `gorm:"size:100;null;" json:"lainnya"`
	UpdatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	TransactionsID   uint64    `json:"transactions_id,omitempty" gorm:"null"`
}

//Nonairs : Struct list Nonair
type Nonairs []Nonair

//PeriodeNonair : Struct Summary Periode NonAir
type PeriodeNonair struct {
	Periode         string  `gorm:"size:100;not null;" json:"periode"`
	Nosamb          string  `gorm:"size:100;not null;" json:"nosamb"`
	Bulan           string  `gorm:"size:100;not null;" json:"bulan"`
	Status          string  `gorm:"size:100;not null;" json:"status"`
	TotalPerPeriode float64 `gorm:"size:100;null;" json:"total_per_periode"`
}

//PeriodeNonairs : Struct list PeriodeNonair
type PeriodeNonairs []PeriodeNonair

//AngsuranNonAir : angsuran Non Air
type AngsuranNonAir struct {
	Nomor          string  `json:"nomor"`
	Periode        string  `json:"periode,omitempty"`
	Jenis          string  `json:"jenis"`
	TotalTagihan   float64 `json:"total_tagihan"`
	Administrasi   float64 `json:"administrasi"`
	Total          float64 `json:"total"`
	DendaTunggakan float64 `json:"denda_tunggakan"`
	JumlahTermin   string  `json:"jumlah_termin"`
	Keterangan     string  `json:"keterangan"`
	SisaTagihan    float64 `json:"sisa_tagihan"`
}

//AngsuranNonAirs  : angsuran Non Air
type AngsuranNonAirs []AngsuranNonAir
