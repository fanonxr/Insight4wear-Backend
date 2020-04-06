package models

import (
	"github.com/globalsign/mgo/bson"
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




