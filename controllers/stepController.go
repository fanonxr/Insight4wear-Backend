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
type StepController struct {}

// functions to create collections for each sensor
func StepCollection(c *mongo.Database) {
	stepCollection = c.Collection("StepData")
}

func GetAllStepData(c *gin.Context) {
	var data []models.StepSensorData

	c.BindJSON(&data)

	cursor, err := stepCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while fetching all Activity sensor data, Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	}

	// iterate through the returned cursor
	for cursor.Next(context.TODO()) {
		var stepData models.StepSensorData
		cursor.Decode(&stepData)
		// add the activity data to the list
		data = append(data, stepData)
	}
	// Bind the retrieve data from the request
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": data,
	})
}

// method to create a activity data within mongodb
func CreateStepData(c *gin.Context) {
	var data models.StepSensorData

	c.BindJSON(&data)

	stepCountDeleta := data.StepCountDelta
	startTime := data.StartTime
	endTime := data.EndTime

	buildStepData := models.StepSensorData{
		TimeStamp: models.TimeStamp{
			StartTime: startTime,
			EndTime: endTime,
		},
		StepCountDelta: stepCountDeleta,
	}

	_, err := stepCollection.InsertOne(context.TODO(), buildStepData)

	if err != nil {
		log.Printf("Error while inserting activity data into the db, Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"message": "Step Data created successfully",
		"data": buildStepData,
	})

}

// get a single activity data
func GetSingleStepData(c *gin.Context) {
	todoId := c.Param("todoId")

	data := models.ActivitySensorData{}
	err := stepCollection.FindOne(context.TODO(), bson.M{"id": todoId}).Decode(&data)
	if err != nil {
		log.Printf("Error while getting a single da, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Todo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Acitivty Data",
		"data":    data,
	})
	return
}

// method to delete a single activity ids
func DeleteStepData(c *gin.Context) {
	dataId := c.Param("id")

	_, err := stepCollection.DeleteOne(context.TODO(), bson.M{"id": dataId})

	if err != nil {
		log.Printf("Error while deleting a single Activity, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Step Data deleted successfully",
	})
	return
}