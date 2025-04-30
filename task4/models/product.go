package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name	string	`json:"name"`
	Price	float64	`json:"price"`
	Quantity	float64	`json:"quantity"`
	SupplierId	uint	`json:"supplier_id"`
	CategoryId	uint	`json:"category_id"`
	Category	Category	`gorm:"foreignKey:CategoryId"`
	Supplier	Supplier	`gorm:"foreignKey:SupplierId"`
}