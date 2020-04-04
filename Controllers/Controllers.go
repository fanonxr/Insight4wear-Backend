package Controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"insight4wear-backend/Models"
	"log"
	"net/http"
)

// collection for the db
var collection *mongo.Collection

// functions to create collections for each sensor
func ActivityCollection(c *mongo.Database) {
	collection = c.Collection("ActivityData")
}

func GetAllActivityData(c *gin.Context) {
	var data []Models.ActivitySensorData

	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while fetching all Activity sensor data, Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	}

	// iterate through the returned cursor
	for cursor.Next(context.TODO()) {
		var activityData Models.ActivitySensorData
		cursor.Decode(&activityData)
		data = append(data, activityData) // add the activity data to the list
	}
}

// method to create a activity data within mongodb
func CreateActivityData(c *gin.Context) {
	var data Models.ActivitySensorData

	c.BindJSON(&data)

	// get all data out of object
	// id := uuid.Must(uuid.NewV4())
	duration := data.Duration
	activity := data.Activity
	startTime := data.StartTime
	endTime := data.EndTime

	buildActivityData := Models.ActivitySensorData{
		CommonSensorData: Models.CommonSensorData{
			StartTime: startTime,
			EndTime: endTime,
		},
		Activity:         activity,
		Duration:         duration,
	}

	_, err := collection.InsertOne(context.TODO(), buildActivityData)

	if err != nil {
		log.Printf("Error while inserting activity data into the db, Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"message": "Activity Data created successfully",
	})

}

// get a single activity data
func GetSingleActivityData(c *gin.Context) {
	todoId := c.Param("todoId")

	data := Models.ActivitySensorData{}
	err := collection.FindOne(context.TODO(), bson.M{"id": todoId}).Decode(&data)
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
func DeleteTodo(c *gin.Context) {
	dataId := c.Param("id")

	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": dataId})
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
		"message": "Todo deleted successfully",
	})
	return
}
