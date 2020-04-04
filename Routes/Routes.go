package Routes

import (
	"github.com/gin-gonic/gin"
	"insight4wear-backend/Controllers"
	"net/http"
)

// function to setup Routes for the REST API
func SetupRouter() *gin.Engine {
	router := gin.Default()
	appApi := router.Group("/app")
	{
		// TODO: Add other sensor Routes once route for this is working
		appApi.GET("sensor") // Route to display all sensor data

		appApi.GET("sensor/activity", Controllers.GetAllActivityData) // route to display the activity sensor data
		appApi.GET("sensor/activity/:id", Controllers.GetSingleActivityData) // route to display the 1 specific activity data
		appApi.POST("sensor/activity/:id", Controllers.CreateActivityData) // route to post the activity data
		appApi.DELETE("sensor/activity", Controllers.DeleteTodo) // route to delete the activity data
	}

	return router
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To the Insight4Wear API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}



