package models

import (
	"github.com/globalsign/mgo/bson"
)

// structure of activity sensor data
type StepSensorData struct {
	ID bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Timestamp TimeStamp `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	StepCountDelta int `json:"step_count_delta,omitempty" bson:"step_count_delta,omitempty"`
}