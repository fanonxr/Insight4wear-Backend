package routes

import (
	"github.com/gin-gonic/gin"
	"insight4wear-backend/controllers"
	"net/http"
)

// function to setup routes for the REST API
func SetupRouter() *gin.Engine {
	// create router with Gin
	router := gin.Default()
	// Handle welcome route
	router.GET("/", welcome)

	v1 := router.Group("/api/v1")
	{
		// TODO: Add other sensor routes once route for this is working
		v1.GET("sensor") // Route to display all sensor data

		v1.GET("sensor/activity", controllers.GetAllActivityData)    // route to display the activity sensor data
		v1.POST("sensor/activity", controllers.CreateActivityData)   // route to post the activity data
		v1.DELETE("sensor/activity", controllers.DeleteActivityData) // route to delete the activity data
	}

	// Handle error response when a route is not defined
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"status":  404,
			"message": "Route Not Found",
		})
	})

	return router
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To the Insight4Wear API",
	})
	return
}




