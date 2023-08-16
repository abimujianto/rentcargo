package models

type Merk struct {
	ID   int    `gorm:"primaryKey:autoIncrement" json:"id"`
	Merk string `json:"merk"`
}

var Merks []Merk
