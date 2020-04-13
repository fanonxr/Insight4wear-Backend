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
type PowerController struct {}

// functions to create collections for each sensor
func PowerCollection(c *mongo.Database) {
	powerCollection = c.Collection("PowerData")
}

func GetAllPowerData(c *gin.Context) {
	var data []models.PowerSensorData

	c.BindJSON(&data)

	cursor, err := powerCollection.Find(context.TODO(), bson.M{})

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
		var powerData models.PowerSensorData
		cursor.Decode(&powerData)
		// add the activity data to the list
		data = append(data, powerData)
	}
	// Bind the retrieve data from the request
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Power Data created successfully",
		"data": data,
	})
	return
}

// method to create a activity data within mongodb
func CreatePowerData(c *gin.Context) {
	var data models.PowerSensorData

	c.BindJSON(&data)

	watts := data.Watts
	startTime := data.StartTime
	endTime := data.EndTime

	buildPowerData := models.PowerSensorData{
		TimeStamp: models.TimeStamp{
			StartTime: startTime,
			EndTime: endTime,
		},
		Watts: watts,
	}

	_, err := powerCollection.InsertOne(context.TODO(), buildPowerData)

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
		"message": "Power dara successfully created",
		"data": buildPowerData,
	})
	return
}

// method to delete a single activity ids
func DeletePowerData(c *gin.Context) {
	dataId := c.Param("id")

	_, err := powerCollection.DeleteOne(context.TODO(), bson.M{"id": dataId})

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
		"message": "Power Data deleted successfully",
	})
	return
}
