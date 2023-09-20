package controllers

import (
	"net/http"

	"github.com/Godofin/travel-planner/database"
	"github.com/Godofin/travel-planner/models"
	"github.com/gin-gonic/gin"
)

// Passo a passo de gerar autorização:

// 1 - Criar endpoint de login e senha
// 2 - Gerar token JWT
// 3 - Middleware de Autenticação

// List general database data
func ReadUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(200, users)
}

// List user by ID
func ReadUserById(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	database.DB.First(&user, id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "User not found.",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Register data in database
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erros": err.Error(),
		})
		return
	}
	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

// Delete user by ID
func DeleteById(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	database.DB.Delete(&user, id)
	c.JSON(http.StatusOK, gin.H{"data": "User successfully deleted"})
}

func EditUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	database.DB.First(&user, id)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Model(&user).UpdateColumns(user)
	c.JSON(http.StatusOK, user)
}
