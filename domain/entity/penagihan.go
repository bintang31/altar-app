package entity

//Penagihan : Struct Entity Penagihan
type Penagihan struct {
	Nosamb         string  `gorm:"size:100;not null;" json:"nosamb"`
	Nama           string  `gorm:"size:255;not null;" json:"nama"`
	Alamat         string  `gorm:"size:255;null;" json:"alamat"`
	Notelp         string  `gorm:"size:100;null;" json:"notelp"`
	KodePdam       string  `gorm:"size:10;null;" json:"kode_pdam"`
	Pdam           string  `gorm:"size:10;null;" json:"pdam"`
	RayonName      string  `gorm:"size:10;null;" json:"rayon_name"`
	Kolektif       string  `gorm:"size:100;null;" json:"kolektif"`
	StatusKolektif string  `gorm:"size:100;null;" json:"status_kolektif"`
	TagihanAir     int     `gorm:"size:100;null;" json:"tagihan_air"`
	TotalTagihan   float64 `gorm:"type:numeric;null;" json:"total_tagihan"`
}

//Penagihans : Struct list Penagihan
type Penagihans []Penagihan

//PenagihansSrKolektif : Struct Entity Penagihan
type PenagihansSrKolektif struct {
	Nosamb             string  `gorm:"size:100;not null;" json:"nosamb"`
	Nama               string  `gorm:"size:255;not null;" json:"nama"`
	Alamat             string  `gorm:"size:255;null;" json:"alamat"`
	Notelp             string  `gorm:"size:100;null;" json:"notelp"`
	KodePdam           string  `gorm:"size:10;null;" json:"kode_pdam"`
	Pdam               string  `gorm:"size:10;null;" json:"pdam"`
	RayonName          string  `gorm:"size:10;null;" json:"rayon_name"`
	Kolektif           string  `gorm:"size:100;null;" json:"kolektif"`
	StatusKolektif     string  `gorm:"size:100;null;" json:"status_kolektif"`
	StatusPelanggan    string  `gorm:"size:100;null;" json:"status_pelanggan"`
	StatusBilling      string  `gorm:"size:100;null;" json:"status_billing"`
	TagihanAir         int     `gorm:"size:100;null;" json:"tagihan_air"`
	TotalTagihanAir    float64 `gorm:"type:numeric;null;" json:"total_tagihan_air"`
	TagihanNonAir      int     `gorm:"size:100;null;" json:"tagihan_nonair"`
	TotalTagihanNonAir float64 `gorm:"type:numeric;null;" json:"total_tagihan_nonair"`
	TotalDenda         float64 `gorm:"type:numeric;null;" json:"total_denda"`
	TotalAdministrasi  float64 `gorm:"type:numeric;null;" json:"total_administrasi"`
	TotalTagihan       float64 `gorm:"type:numeric;null;" json:"total_tagihan"`
	PeriodeTagihan     string  `gorm:"size:100;null;" json:"periode_tagihan"`
	PeriodeKolektif    string  `gorm:"size:100;null;" json:"periode_kolektif"`
}

//PenagihansSrKolektifs : Struct list PenagihanSrKolektifs
type PenagihansSrKolektifs []PenagihansSrKolektif

//PenagihansParams : Struct Penagihan With Parameter
type PenagihansParams struct {
	Filter string `form:"filter" json:"filter"`
	Page   int    `form:"page" json:"page"`
	UserID uint64 `form:"user" json:"user"`
}

//PenagihanBermasalah : Struct Entity PenagihanBermasalah
type PenagihanBermasalah struct {
	Nosamb string `gorm:"size:100;not null;" json:"nosamb"`
	Notes  string `gorm:"size:255;not null;" json:"notes"`
}

//PenagihanBermasalahs : Struct list Penagihan
type PenagihanBermasalahs []PenagihanBermasalah

//ResponseLoket : Response From Loket Module
type ResponseLoket struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

//InquiryData : from inquiry endpoint
type InquiryData struct {
	Data     map[string]interface{} `json:"data"`
	Response map[string]interface{} `json:"response"`
}
