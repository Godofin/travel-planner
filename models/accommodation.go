package models

import (
	"gorm.io/gorm"
)

type Accommodation struct {
	gorm.Model
	Name          string  `gorm:"type:varchar(255)" json:"name"`
	Description   string  `gorm:"type:varchar(255)" json:"description"`
	Latitude      float64 `gorm:"type:decimal(10,8)" json:"latitude"`
	Longitude     float64 `gorm:"type:decimal(10,8)" json:"longitude"`
	Type          string  `gorm:"type:varchar(255)" json:"type"`
	Photo         string  `gorm:"type:varchar(255)" json:"photo"`
	DestinationID int     `json:"destination_id"`
}
