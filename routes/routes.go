package routes

import (
	controllers "github.com/Godofin/travel-planner/controller"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/users", controllers.ReadUsers)
	r.GET("/users/:id", controllers.ReadUserById)
	r.POST("/users", controllers.CreateUser)
	r.DELETE("/users/:id", controllers.DeleteById)
	r.PATCH("/users/:id", controllers.EditUser)

	r.Run()
}
