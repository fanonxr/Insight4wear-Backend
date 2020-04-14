package models

import "github.com/globalsign/mgo/bson"

// structure of Calorie sensor data
type CalorieSensorData struct {
	ID bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Timestamp TimeStamp `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	Calories float64 `json:"calories,omitempty" bson:"calories,omitempty"`
}
