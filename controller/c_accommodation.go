package controllers

import (
	"net/http"

	"github.com/Godofin/travel-planner/database"
	"github.com/Godofin/travel-planner/models"
	"github.com/gin-gonic/gin"
)

// List general database data
func ReadAccommodation(c *gin.Context) {
	var accommodations []models.Accommodation
	database.DB.Find(&accommodations)
	c.JSON(200, accommodations)
}

// List Accommodation by ID
func ReadAccommodationById(c *gin.Context) {
	var accommodations models.Accommodation
	id := c.Params.ByName("id")
	database.DB.First(&accommodations, id)
	if accommodations.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Accommodation not found.",
		})
		return
	}
	c.JSON(http.StatusOK, accommodations)
}

// Register data in database
func CreateAccommodation(c *gin.Context) {
	var accommodations models.Accommodation
	if err := c.ShouldBindJSON(&accommodations); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erros": err.Error(),
		})
		return
	}
	database.DB.Create(&accommodations)
	c.JSON(http.StatusOK, accommodations)
}

// Delete Accommodation by ID
func DeleteByIdAccommodation(c *gin.Context) {
	var accommodations models.Accommodation
	id := c.Params.ByName("id")
	database.DB.Delete(&accommodations, id)
	c.JSON(http.StatusOK, gin.H{"data": "Accommodation successfully deleted"})
}

func EditAccommodation(c *gin.Context) {
	var accommodations models.Accommodation
	id := c.Params.ByName("id")
	database.DB.First(&accommodations, id)
	if err := c.ShouldBindJSON(&accommodations); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Model(&accommodations).UpdateColumns(accommodations)
	c.JSON(http.StatusOK, accommodations)
}
