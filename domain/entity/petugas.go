package entity

//Petugas : Struct Entity Petugas
type Petugas struct {
	ID          uint64 `gorm:"primary_key" json:"id"`
	NamaPetugas string `gorm:"size:255;not null;" json:"nama_petugas"`
}

//PetugasData : Struct Entity Petugas Data
type PetugasData struct {
	PenagihanPelanggan []PenagihansSrKolektif `json:"penagihan_pelanggan"`
	PenagihanBilling   []Drd                  `json:"penagihan_billing"`
}
