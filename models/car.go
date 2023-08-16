package models

type Car struct {
	ID          int     `gorm:"primaryKey:autoIncrement" json:"id"`
	IDMerk      int     `gorm:"foreignKey:IDMerk" json:"id_merk"`
	Merk        Merk    `gorm:"foreignKey:IDMerk"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

var Cars []Car
