package models

import (
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Name          string  `gorm:"type:varchar(255)" json:"name"`
	Description   string  `gorm:"type:varchar(255)" json:"description"`
	Latitude      float64 `gorm:"type:decimal(10,8)" json:"latitude"`
	Longitude     float64 `gorm:"type:decimal(10,8)" json:"longitude"`
	Photo         string  `gorm:"type:varchar(255)" json:"photo"`
	DestinationID int     `json:"destination_id"`
}
