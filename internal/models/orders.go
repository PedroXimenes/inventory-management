package models

import "gorm.io/gorm"

type Orders struct {
	gorm.Model

	UserId string `json:"user_id"`
	Status string `json:"status"`
}
