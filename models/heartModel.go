package models

import "github.com/globalsign/mgo/bson"

// structure of Heart sensor data
type HeartSensorData struct {
	ID bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	TimeStamp
	BPM float64 `json:"bpm,omitempty" bson:"bpm,omitempty"`
}