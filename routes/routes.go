package routes

import (
	c_accommodations "github.com/Godofin/travel-planner/controller"
	c_activity "github.com/Godofin/travel-planner/controller"
	c_destination "github.com/Godofin/travel-planner/controller"
	c_itinerary "github.com/Godofin/travel-planner/controller"
	c_users "github.com/Godofin/travel-planner/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:5500"}                     // Troque para sua origem necessária, use * para permitir todas
	config.AllowMethods = []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"} // HTTP methods your API supports
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}

	r.Use(cors.New(config))
	//Users
	r.GET("/users", c_users.ReadUsers)
	r.GET("/users/:id", c_users.ReadUserById)
	r.POST("/users", c_users.CreateUser)
	r.DELETE("/users/:id", c_users.DeleteById)
	r.PATCH("/users/:id", c_users.EditUser)
	itineraries := r.Group("/itinerary")
	// Aplicar configurações CORS a este grupo de rotas
	itineraries.Use(cors.New(config))
	{
		itineraries.GET("", c_itinerary.ReadItinerary)
		itineraries.GET("/:id", c_itinerary.ReadItineraryById)
		itineraries.POST("", c_itinerary.CreateItinerary)
		itineraries.DELETE("/:id", c_itinerary.DeleteByIditinerary)
		itineraries.PATCH("/:id", c_itinerary.EditItinerary)
	}
	//Itnerary details
	r.GET("/itnerarydetails", c_itinerary.ReadItineraryDetails)
	r.GET("/itnerarydetails/:id", c_itinerary.ReadItineraryDetailsById)
	r.POST("/itnerarydetails", c_itinerary.CreateItineraryDetails)
	r.DELETE("/itnerarydetails/:id", c_itinerary.DeleteByIditineraryDetails)
	r.PATCH("/itnerarydetails/:id", c_itinerary.EditItineraryDetails)
	//Destination
	r.GET("/destination", c_destination.ReadDestination)
	r.GET("/destination/:id", c_destination.ReadDestinationById)
	r.POST("/destination", c_destination.CreateDestination)
	r.DELETE("/destination/:id", c_destination.DeleteByIdDestination)
	r.PATCH("/destination/:id", c_destination.EditDestination)
	//Activity
	r.GET("/activity", c_activity.ReadActivity)
	r.GET("/activity/:id", c_activity.ReadActivityById)
	r.POST("/activity", c_activity.CreateActivity)
	r.DELETE("/activity/:id", c_activity.DeleteByIdActivity)
	r.PATCH("/activity/:id", c_activity.EditActivity)
	//accommodations
	r.GET("/accommodations", c_accommodations.ReadAccommodation)
	r.GET("/accommodations/:id", c_accommodations.ReadAccommodationById)
	r.POST("/accommodations", c_accommodations.CreateAccommodation)
	r.DELETE("/accommodations/:id", c_accommodations.DeleteByIdAccommodation)
	r.PATCH("/accommodations/:id", c_accommodations.EditAccommodation)

	r.Run()
}
