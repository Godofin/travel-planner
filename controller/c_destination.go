package controllers

import (
	"net/http"

	"github.com/Godofin/travel-planner/database"
	"github.com/Godofin/travel-planner/models"
	"github.com/gin-gonic/gin"
)

// List general database data
func ReadDestination(c *gin.Context) {
	var destinations []models.Destination
	database.DB.Find(&destinations)
	c.JSON(200, destinations)
}

// List Destination by ID
func ReadDestinationById(c *gin.Context) {
	var destinations models.Destination
	id := c.Params.ByName("id")
	database.DB.First(&destinations, id)
	if destinations.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Destination not found.",
		})
		return
	}
	c.JSON(http.StatusOK, destinations)
}

// Register data in database
func CreateDestination(c *gin.Context) {
	var destinations models.Destination
	if err := c.ShouldBindJSON(&destinations); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erros": err.Error(),
		})
		return
	}
	database.DB.Create(&destinations)
	c.JSON(http.StatusOK, destinations)
}

// Delete Destination by ID
func DeleteByIdDestination(c *gin.Context) {
	var destinations models.Destination
	id := c.Params.ByName("id")
	database.DB.Delete(&destinations, id)
	c.JSON(http.StatusOK, gin.H{"data": "Destination successfully deleted"})
}

func EditDestination(c *gin.Context) {
	var destinations models.Destination
	id := c.Params.ByName("id")
	database.DB.First(&destinations, id)
	if err := c.ShouldBindJSON(&destinations); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Model(&destinations).UpdateColumns(destinations)
	c.JSON(http.StatusOK, destinations)
}
