package models

import (
	"time"

	"gorm.io/gorm"
)

type ItineraryDetails struct {
	gorm.Model
	ItineraryID     int       `json:"itinerary_id"`
	ActivityID      int       `json:"activity_id"`
	AccommodationID int       `json:"accommodation_id"`
	Date            time.Time `json:"date"`
	Notes           string    `gorm:"type:varchar(255)" json:"notes"`
}
