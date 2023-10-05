package controllers

import (
	"net/http"

	"github.com/Godofin/travel-planner/database"
	"github.com/Godofin/travel-planner/models"
	"github.com/gin-gonic/gin"
)

// List general database data
func ReadItinerary(c *gin.Context) {
	var itinerarys []models.Itinerary
	database.DB.Find(&itinerarys)
	c.JSON(200, itinerarys)
}

// List Itinerary by ID
func ReadItineraryById(c *gin.Context) {
	var itinerarys models.Itinerary
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
func CreateItinerary(c *gin.Context) {
	var itinerarys models.Itinerary
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
func DeleteByIditinerary(c *gin.Context) {
	var itinerarys models.Itinerary
	id := c.Params.ByName("id")
	database.DB.Delete(&itinerarys, id)
	c.JSON(http.StatusOK, gin.H{"data": "Itinerary successfully deleted"})
}

func EditItinerary(c *gin.Context) {
	var itinerarys models.Itinerary
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
