package models

import (
	"time"

	"gorm.io/gorm"
)

type Itinerary struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Start_Date  time.Time `json:"start_date"`
	End_Date    time.Time `json:"end_date"`
	User_Id     string    `json:"user_id"`
}
