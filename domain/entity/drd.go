package entity

//Drd : Struct Entity Drd
type Drd struct {
	Nosamb         string  `json:"nosamb" gorm:"primary_key"`
	Periode        string  `json:"periode" gorm:"not null; unique; size:50"`
	Pdam           string  `json:"pdam" gorm:"not null; unique; size:50"`
	Pakai          int     `json:"pakai" gorm:"null"`
	Stanlalu       int     `json:"stanlalu" gorm:"null"`
	Stanangkat     int     `json:"stanangkat" gorm:"null"`
	Stanskrg       int     `json:"stanskrg" gorm:"null"`
	Danameter      int     `json:"danameter" gorm:"null"`
	AirLimbah      int     `json:"airlimbah" gorm:"null"`
	Biayapemakaian float64 `json:"biayapemakaian" gorm:"null"`
	Rekair         float64 `json:"rekair" gorm:"null"`
	Administrasi   float64 `json:"administrasi" gorm:"null"`
	Pemeliharaan   float64 `json:"pemeliharaan" gorm:"null"`
	Retribusi      float64 `json:"retribusi" gorm:"null"`
	LLTT           float64 `json:"lltt" gorm:"null"`
	Materai        float64 `json:"materai" gorm:"null"`
	PPN            float64 `json:"ppn" gorm:"null"`
	Dendatunggakan float64 `json:"dendatunggakan" gorm:"null"`
	Segel          float64 `json:"segel" gorm:"null"`
	Bulan          string  `json:"bulan" gorm:"null; size:10"`
	Kolektif       string  `json:"kolektif" gorm:"null; size:10"`
	Angsur         string  `json:"angsur" gorm:"null; size:10"`
	Batal          string  `json:"batal" gorm:"null; size:10"`
	Lunas          string  `json:"lunas" gorm:"null; size:10"`
	Total          float64 `json:"total" gorm:"null"`
	Tglupload      int64   `json:"tglupload"`
	Tglbayar       int64   `json:"tglbayar"`
	TglbayarString string  `json:"tglbayar_string" gorm:"null"`
	TransactionsID int     `json:"transactions_id" gorm:"null"`
}

//Drds : Struct list DRD
type Drds []Drd
