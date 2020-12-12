package entity

//Penagihan : Struct Entity Penagihan
type Penagihan struct {
	Nosamb string `gorm:"size:100;not null;" json:"nosamb"`
	Nama   string `gorm:"size:255;not null;" json:"nama"`
	Alamat string `gorm:"size:255;null;" json:"alamat"`
}

//Penagihans : Struct list Penagihan
type Penagihans []Penagihan

//PenagihansSrKolektif : Struct Entity Penagihan
type PenagihansSrKolektif struct {
	Nosamb       string  `gorm:"size:100;not null;" json:"nosamb"`
	Nama         string  `gorm:"size:255;not null;" json:"nama"`
	Alamat       string  `gorm:"size:255;null;" json:"alamat"`
	TotalTagihan float64 `gorm:"type:numeric;null;" json:"total_tagihan"`
}

//PenagihansSrKolektifs : Struct list PenagihanSrKolektifs
type PenagihansSrKolektifs []PenagihansSrKolektif

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
