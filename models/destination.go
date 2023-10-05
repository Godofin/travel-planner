package models

import "gorm.io/gorm"

type Destination struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Photo       string `json:"photo"`
}
