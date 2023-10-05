package controllers

import (
	"net/http"

	"github.com/Godofin/travel-planner/database"
	"github.com/Godofin/travel-planner/models"
	"github.com/gin-gonic/gin"
)

// List general database data
func ReadActivity(c *gin.Context) {
	var activitys []models.Activity
	database.DB.Find(&activitys)
	c.JSON(200, activitys)
}

// List Activity by ID
func ReadActivityById(c *gin.Context) {
	var activitys models.Activity
	id := c.Params.ByName("id")
	database.DB.First(&activitys, id)
	if activitys.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Activity not found.",
		})
		return
	}
	c.JSON(http.StatusOK, activitys)
}

// Register data in database
func CreateActivity(c *gin.Context) {
	var activitys models.Activity
	if err := c.ShouldBindJSON(&activitys); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erros": err.Error(),
		})
		return
	}
	database.DB.Create(&activitys)
	c.JSON(http.StatusOK, activitys)
}

// Delete Activity by ID
func DeleteByIdActivity(c *gin.Context) {
	var activitys models.Activity
	id := c.Params.ByName("id")
	database.DB.Delete(&activitys, id)
	c.JSON(http.StatusOK, gin.H{"data": "Activity successfully deleted"})
}

func EditActivity(c *gin.Context) {
	var activitys models.Activity
	id := c.Params.ByName("id")
	database.DB.First(&activitys, id)
	if err := c.ShouldBindJSON(&activitys); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Model(&activitys).UpdateColumns(activitys)
	c.JSON(http.StatusOK, activitys)
}
