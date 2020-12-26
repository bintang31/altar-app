package entity

//Nonair : Struct Entity Nonair
type Nonair struct {
	Nomor          string  `gorm:"size:100;not null;" json:"nomor"`
	Periode        string  `gorm:"size:100;not null;" json:"periode"`
	Jenis          string  `gorm:"size:100;not null;" json:"jenis,omitempty"`
	JenisTagihan   string  `gorm:"size:100;not null;" json:"jenis_tagihan"`
	Lunas          string  `gorm:"size:100;null;" json:"lunas"`
	Total          float64 `gorm:"size:100;null;" json:"total"`
	Administrasi   float64 `gorm:"size:100;null;" json:"administrasi"`
	Biayapasang    float64 `gorm:"size:100;null;" json:"biayapasang"`
	DendaTunggakan float64 `gorm:"size:100;null;" json:"denda_tunggakan"`
	Lainnya        float64 `gorm:"size:100;null;" json:"lainnya"`
	TransactionsID uint64  `json:"transactions_id,omitempty" gorm:"null"`
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
