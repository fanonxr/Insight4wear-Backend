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
		v1.GET("sensor") // Route to display all sensor data

		// routes for activity data
		v1.GET("sensor/activity", controllers.GetAllActivityData)
		v1.POST("sensor/activity", controllers.CreateActivityData)
		v1.DELETE("sensor/activity", controllers.DeleteActivityData)

		// routes for calorie data
		v1.GET("sensor/calorie", controllers.GetAllCalorieData)
		v1.POST("sensor/calorie", controllers.CreateCalorieData)
		v1.DELETE("sensor/calorie", controllers.DeleteCalorieData)

		// routes for Heart data
		v1.GET("sensor/heart", controllers.GetAllHeartData)
		v1.POST("sensor/heart", controllers.CreateHeartData)
		v1.DELETE("sensor/heart", controllers.DeleteHeartData)

		// routes for Power data
		v1.GET("sensor/power", controllers.GetAllPowerData)
		v1.POST("sensor/power", controllers.CreatePowerData)
		v1.DELETE("sensor/power", controllers.DeletePowerData)

		// routes for Step Data
		v1.GET("sensor/steps", controllers.GetAllStepData)
		v1.POST("sensor/steps", controllers.CreateStepData)
		v1.DELETE("sensor/steps", controllers.DeleteStepData)
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




