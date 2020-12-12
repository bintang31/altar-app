package entity

//Nonair : Struct Entity Nonair
type Nonair struct {
	Nomor   string  `gorm:"size:100;not null;" json:"nomor"`
	Periode string  `gorm:"size:100;not null;" json:"periode"`
	Jenis   string  `gorm:"size:100;not null;" json:"jenis"`
	Total   float64 `gorm:"size:100;null;" json:"total"`
}

//Nonairs : Struct list Nonair
type Nonairs []Nonair
