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
type HeartController struct {}

// functions to create collections for each sensor
func HeartCollection(c *mongo.Database) {
	heartCollection = c.Collection("HeartData")
}

func GetAllHeartData(c *gin.Context) {
	var data []models.HeartSensorData

	c.BindJSON(&data)

	cursor, err := heartCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while fetching all Calorie sensor data, Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	}

	// iterate through the returned cursor
	for cursor.Next(context.TODO()) {
		var heartData models.HeartSensorData
		cursor.Decode(&heartData)
		// add the activity data to the list
		data = append(data, heartData)
	}
	// Bind the retrieve data from the request
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Heart Data created successfully",
		"data": data,
	})
}

// method to create a activity data within mongodb
func CreateHeartData(c *gin.Context) {
	var data models.HeartSensorData

	c.BindJSON(&data)

	BPM := data.BPM
	startTime := data.StartTime
	endTime := data.EndTime

	buildHeartData := models.HeartSensorData{
		TimeStamp: models.TimeStamp{
			StartTime: startTime,
			EndTime: endTime,
		},
		BPM: BPM,
	}

	_, err := heartCollection.InsertOne(context.TODO(), buildHeartData)

	if err != nil {
		log.Printf("Error while inserting Calorie data into the db, Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":       http.StatusCreated,
		"data": buildHeartData,
	})
}

// method to delete a single activity ids
func DeleteHeartData(c *gin.Context) {
	dataId := c.Param("id")

	_, err := heartCollection.DeleteOne(context.TODO(), bson.M{"id": dataId})

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
		"message": "Heart Data deleted successfully",
	})
	return
}
