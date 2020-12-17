package entity

//Pelanggan : Struct Entity Pelanggan
type Pelanggan struct {
	Nosamb   string `gorm:"size:100;not null;" json:"nosamb"`
	Nama     string `gorm:"size:255;not null;" json:"nama"`
	Alamat   string `gorm:"size:255;null;" json:"alamat"`
	Notelp   string `gorm:"size:50;not null;" json:"notelp"`
	Golongan string `gorm:"size:50;not null;" json:"golongan"`
}

//Pelanggans : Struct list Pelanggan
type Pelanggans []Pelanggan

//InputInquiryPelanggan : Struct Entity InputInquiryPelanggan
type InputInquiryPelanggan struct {
	Nosamb string `json:"nosamb"`
	Pdam   string `json:"pdam"`
}

//InquiryCollection : Struct Entity InquiryCollection
type InquiryCollection struct {
	Rekair       []RekairDetail `json:"rekair"`
	Totaltagihan int            `json:"totaltagihan"`
}

//RekairDetail : Struct Inquiry rekening air
type RekairDetail struct {
	Alamat         string  `json:"alamat"`
	Administrasi   float64 `json:"administrasi"`
	Periode        string  `json:"periode"`
	Denda          float64 `json:"denda"`
	Retribusi      float64 `json:"retribusi"`
	BiayaPemakaian float64 `json:"biayapemakaian"`
	Tagihan        float64 `json:"tagihan"`
	Bulan          string  `json:"bulan"`
	Nama           string  `json:"nama"`
	Nosamb         string  `json:"nosamb"`
}

//RekairDetails : Struct list Pelanggan
type RekairDetails []RekairDetail
