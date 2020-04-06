package models

import (
	"github.com/globalsign/mgo/bson"
)

// structure of activity sensor data
type ActivitySensorData struct {
	ID bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	TimeStamp
	Activity int `json:"activity,omitempty" bson:"activity,omitempty"`
	Duration float32 `json:"duration,omitempty" bson:"duration,omitempty"`
}




