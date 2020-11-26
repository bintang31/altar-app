package entity

//Pelanggan : Struct Entity Pelanggan
type Pelanggan struct {
	Nosamb string `gorm:"size:100;not null;" json:"nosamb"`
	Nama   string `gorm:"size:255;not null;" json:"nama"`
	Alamat string `gorm:"size:255;null;" json:"alamat"`
	Notelp string `gorm:"size:50;not null;" json:"notelp"`
}
