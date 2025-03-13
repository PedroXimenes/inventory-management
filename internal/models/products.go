package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model

	Name  string  `json:"name"`
	Sku   string  `json:"sku"`
	Price float64 `json:"price"`
}
