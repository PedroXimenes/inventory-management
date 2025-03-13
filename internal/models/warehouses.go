package models

type Warehouses struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Location string `json:"location"`
}
