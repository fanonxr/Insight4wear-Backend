package models

import (
	"github.com/globalsign/mgo/bson"
	"log"
)
// SensorData interface that each type of sensor model will implement to interact with the database
type SensorData interface {
	unMap()
}

// representing the shared values that each sensor will have
type TimeStamp struct {
	StartTime string `json:"starttime,omitempty" bson:"starttime,omitempty"`
	EndTime string `json:"endtime,omitempty" bson:"endtime,omitempty"`
}

// structure of activity sensor data
type ActivitySensorData struct {
	ID bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	TimeStamp
	Activity int `json:"activity,omitempty" bson:"activity,omitempty"`
	Duration float32 `json:"duration,omitempty" bson:"duration,omitempty"`
}
// How the activity model will be structured
type ActivityModel struct {}

// handle inserting data to the db
func (activityData *ActivityModel) InsertActivityData(data ActivitySensorData) error {
	// connect to the activity connection
	collection := dbConnect.Use(databaseName, "activity")
	// assign result to error object while saving activity data
	err := collection.Insert(bson.M{
		"starttime": data.StartTime,
		"endtime": data.EndTime,
		"activity": data.Activity,
		"duration": data.Duration,
	})

	return err
}

// Handle getting all documents from activity collection
func (activityData *ActivityModel) FetchAllActivityData() []ActivitySensorData {
	var result []ActivitySensorData
	// connect to the activity collection
	collection := dbConnect.Use(databaseName, "activity")

	// fetch all the documents within that database
	err := collection.Find(nil).All(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result
}



