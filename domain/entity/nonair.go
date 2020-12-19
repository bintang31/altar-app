package entity

//Nonair : Struct Entity Nonair
type Nonair struct {
	Nomor          string  `gorm:"size:100;not null;" json:"nomor"`
	Periode        string  `gorm:"size:100;not null;" json:"periode"`
	Jenis          string  `gorm:"size:100;not null;" json:"jenis"`
	Lunas          string  `gorm:"size:100;null;" json:"lunas"`
	Total          float64 `gorm:"size:100;null;" json:"total"`
	TransactionsID uint64  `json:"transactions_id" gorm:"null"`
}

//Nonairs : Struct list Nonair
type Nonairs []Nonair
