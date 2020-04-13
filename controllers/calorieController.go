package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"insight4wear-backend/models"
	"log"
	"net/http"
)

// struct to to handle database operations
type CalorieController struct {}

// functions to create collections for each sensor
func CalorieCollection(c *mongo.Database) {
	calorieCollection = c.Collection("CalorieData")
}

func GetAllCalorieData(c *gin.Context) {
	var data []models.CalorieSensorData

	c.BindJSON(&data)

	cursor, err := calorieCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while fetching all Calorie sensor data, Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	// iterate through the returned cursor
	for cursor.Next(context.TODO()) {
		var calorieData models.CalorieSensorData
		cursor.Decode(&calorieData)
		// add the activity data to the list
		data = append(data, calorieData)
	}
	// Bind the retrieve data from the request
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": data,
	})
	return
}

// method to create a activity data within mongodb
func CreateCalorieData(c *gin.Context) {
	var data models.CalorieSensorData

	c.BindJSON(&data)

	calories := data.Calories
	startTime := data.StartTime
	endTime := data.EndTime

	buildCalorieData := models.CalorieSensorData{
		TimeStamp: models.TimeStamp{
			StartTime: startTime,
			EndTime: endTime,
		},
		Calories: calories,
	}

	_, err := calorieCollection.InsertOne(context.TODO(), buildCalorieData)

	if err != nil {
		log.Printf("Error while inserting Calorie data into the db, Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":       http.StatusCreated,
		"message": "Calorie Data created successfully",
		"data": buildCalorieData,
	})
	return
}

// method to delete a single activity ids
func DeleteCalorieData(c *gin.Context) {
	dataId := c.Param("id")

	_, err := calorieCollection.DeleteOne(context.TODO(), bson.M{"id": dataId})

	if err != nil {
		log.Printf("Error while deleting a single calorie data, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Calorie Data deleted successfully",
	})
	return
}