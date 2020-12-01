package entity

//Penagihan : Struct Entity Penagihan
type Penagihan struct {
	Nosamb string `gorm:"size:100;not null;" json:"nosamb"`
	Nama   string `gorm:"size:255;not null;" json:"nama"`
	Alamat string `gorm:"size:255;null;" json:"alamat"`
}

//Penagihans : Struct list Penagihan
type Penagihans []Penagihan

//PenagihanSrKolektif : Struct Entity Penagihan
type PenagihansSrKolektif struct {
	Nosamb       string  `gorm:"size:100;not null;" json:"nosamb"`
	Nama         string  `gorm:"size:255;not null;" json:"nama"`
	Alamat       string  `gorm:"size:255;null;" json:"alamat"`
	TotalTagihan float64 `gorm:"type:numeric;null;" json:"total_tagihan"`
}

//PenagihanSrKolektifs : Struct list PenagihanSrKolektifs
type PenagihansSrKolektifs []PenagihansSrKolektif
