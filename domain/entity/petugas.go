package entity

//Petugas : Struct Entity Petugas
type Petugas struct {
	ID                     uint64  `gorm:"primary_key" json:"id"`
	NamaPetugas            string  `gorm:"size:255;not null;" json:"nama_petugas"`
	AlamatPdam             string  `gorm:"size:255;null;" json:"alamat_pdam"`
	PhotoBase64            string  `json:"photo_base64" gorm:"not null; size:50"`
	Transaction            int     `json:"transaction" gorm:"not null"`
	TotalTransaction       float64 `json:"total_transaction" gorm:"not null"`
	TotalDiterima          float64 `json:"total_diterima" gorm:"null"`
	LembarTagihanDiterima  int     `json:"lembar_tagihan_diterima" gorm:"not null"`
	SambunganRumahDiterima int     `json:"sambungan_rumah_diterima" gorm:"not null"`
	KolektifDiterima       int     `json:"kolektif_diterima" gorm:"not null"`
	SambunganRumahSetor    int     `json:"sambungan_rumah_setor" gorm:"not null"`
	KolektifSetor          int     `json:"kolektif_setor" gorm:"not null"`
	Setoran                int     `json:"setoran" gorm:"not null"`
	TotalSetoran           float64 `json:"total_setoran" gorm:"null"`
	SisaLimit              float64 `json:"sisa_limit" gorm:"null"`
	LimitPetugas           float64 `json:"limit_petugas" gorm:"null"`
	TglGabung              string  `json:"tgl_gabung" gorm:"null; size:50"`
}

//PetugasData : Struct Entity Petugas Data
type PetugasData struct {
	PenagihanPelanggan       []PenagihansSrKolektif `json:"penagihan_pelanggan"`
	PenagihanBilling         []DrdbyPetugas         `json:"penagihan_billing"`
	PenagihanBillingKolektif []DrdbyPetugas         `json:"penagihan_billing_kolektif"`
	PenagihanBillingNonair   []Nonair               `json:"penagihan_billing_nonair"`
	Petugas                  *Petugas               `json:"petugas"`
}

//Petugass : Struct list Petugas
type Petugass []Petugas
