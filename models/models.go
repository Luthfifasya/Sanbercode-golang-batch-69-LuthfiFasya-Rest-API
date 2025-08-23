package models

type Bioskop struct {
	ID     uint    `gorm:"primaryKey" json:"id"`
	Nama   string  `json:"nama"`
	Lokasi string  `json:"lokasi"`
	Rating float64 `json:"rating"`
}
