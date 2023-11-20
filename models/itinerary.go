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
	Destination string    `json:"destination"`
	Activity    string    `json:"activities"`
	Image       string    `json:"image"`
	User_Id     string    `json:"user_id"`
}
