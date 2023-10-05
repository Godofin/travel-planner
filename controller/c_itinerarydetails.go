package controllers

import (
	"net/http"

	"github.com/Godofin/travel-planner/database"
	"github.com/Godofin/travel-planner/models"
	"github.com/gin-gonic/gin"
)

// List general database data
func ReadItineraryDetails(c *gin.Context) {
	var itinerarys []models.ItineraryDetails
	database.DB.Find(&itinerarys)
	c.JSON(200, itinerarys)
}

// List Itinerary by ID
func ReadItineraryDetailsById(c *gin.Context) {
	var itinerarys models.ItineraryDetails
	id := c.Params.ByName("id")
	database.DB.First(&itinerarys, id)
	if itinerarys.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Itinerary not found.",
		})
		return
	}
	c.JSON(http.StatusOK, itinerarys)
}

// Register data in database
func CreateItineraryDetails(c *gin.Context) {
	var itinerarys models.ItineraryDetails
	if err := c.ShouldBindJSON(&itinerarys); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erros": err.Error(),
		})
		return
	}
	database.DB.Create(&itinerarys)
	c.JSON(http.StatusOK, itinerarys)
}

// Delete Itinerary by ID
func DeleteByIditineraryDetails(c *gin.Context) {
	var itinerarys models.ItineraryDetails
	id := c.Params.ByName("id")
	database.DB.Delete(&itinerarys, id)
	c.JSON(http.StatusOK, gin.H{"data": "Itinerary successfully deleted"})
}

func EditItineraryDetails(c *gin.Context) {
	var itinerarys models.ItineraryDetails
	id := c.Params.ByName("id")
	database.DB.First(&itinerarys, id)
	if err := c.ShouldBindJSON(&itinerarys); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Model(&itinerarys).UpdateColumns(itinerarys)
	c.JSON(http.StatusOK, itinerarys)
}
