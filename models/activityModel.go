package models

import (
	"github.com/globalsign/mgo/bson"
)

// structure of activity sensor data
type ActivitySensorData struct {
	ID bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Timestamp TimeStamp `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	Activity string `json:"activity,omitempty" bson:"activity,omitempty"`
	Duration float32 `json:"duration,omitempty" bson:"duration,omitempty"`
}




