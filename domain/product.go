package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string
	Price    float32
	Discount float32
	Store    string
}
