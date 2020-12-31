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
	Rekair       []RekairDetail   `json:"rekair"`
	Nonair       []NonAirDetail   `json:"nonair"`
	Angsuran     []AngsuranDetail `json:"angsuran"`
	Totaltagihan int              `json:"totaltagihan"`
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
	Action         string  `json:"action"`
	Lunas          string  `json:"lunas"`
}

//RekairDetails : Struct list Pelanggan
type RekairDetails []RekairDetail

//NonAirDetail : Struct Inquiry NonAir
type NonAirDetail struct {
	Kode             string  `json:"kode"`
	Nomor            string  `json:"nomor"`
	Periode          string  `json:"periode"`
	Pdam             string  `json:"pdam"`
	Alamat           string  `json:"alamat"`
	Jenis            string  `json:"jenis"`
	FlagAngsur       int     `json:"flagangsur"`
	Flag             string  `json:"flag"`
	KetJenis         string  `json:"ketjenis"`
	DibebankanKepada string  `json:"dibebankankepada"`
	DendaTunggakan   float64 `json:"dendatunggakan"`
	Total            float64 `json:"total"`
	Administrasi     float64 `json:"administrasi"`
}

//NonAirDetails : Struct list Pelanggan
type NonAirDetails []NonAirDetail

//AngsuranDetail : Struct Inquiry Angsuran
type AngsuranDetail struct {
	Kode             string  `json:"kode"`
	Urutan           string  `json:"urutan"`
	Noangsuran       string  `json:"noangsuran"`
	Nomor            string  `json:"nomor"`
	Periode          string  `json:"periode"`
	Pdam             string  `json:"pdam"`
	Alamat           string  `json:"alamat"`
	Bulan            string  `json:"bulan"`
	Jenis            string  `json:"jenis"`
	FlagAngsur       int     `json:"flagangsur"`
	Termin           int     `json:"termin"`
	Lunas            string  `json:"lunas"`
	Flag             string  `json:"flag"`
	KetJenis         string  `json:"ketjenis"`
	JenisKeterangan  string  `json:"jenis_keterangan"`
	DibebankanKepada string  `json:"dibebankankepada"`
	DendaTunggakan   float64 `json:"dendatunggakan"`
	Total            float64 `json:"total"`
	Administrasi     float64 `json:"administrasi"`
}

//AngsuranDetails : Struct list Pelanggan
type AngsuranDetails []AngsuranDetail

//PeriodeRiwayat : Struct Summary Periode Riwayat Tagihan
type PeriodeRiwayat struct {
	Periode         string  `gorm:"size:100;not null;" json:"periode"`
	Nosamb          string  `gorm:"size:100;not null;" json:"nosamb"`
	Bulan           string  `gorm:"size:100;not null;" json:"bulan"`
	Status          string  `gorm:"size:100;not null;" json:"status"`
	TotalPerPeriode float64 `gorm:"size:100;null;" json:"total_per_periode"`
	TotalPemakaian  float64 `gorm:"size:100;null;" json:"total_pemakaian"`
}

//PeriodeRiwayats : Struct list PeriodeRiwayat
type PeriodeRiwayats []PeriodeRiwayat

//PelangganParams : Struct Pelanggan With Parameter
type PelangganParams struct {
	Pelanggan string `form:"pelanggan" json:"pelanggan"`
}
