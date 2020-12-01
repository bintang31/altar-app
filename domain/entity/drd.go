package entity

//Drd : Struct Entity Drd
type Drd struct {
	Nosamb  string  `gorm:"size:100;not null;" json:"nosamb"`
	Periode string  `gorm:"size:100;not null;" json:"periode"`
	Total   float64 `gorm:"size:100;null;" json:"total"`
}
