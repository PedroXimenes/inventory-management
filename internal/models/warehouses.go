package models

import "gorm.io/gorm"

type Warehouses struct {
	gorm.Model

	Name     string `json:"name"`
	Location string `json:"location"`
}
