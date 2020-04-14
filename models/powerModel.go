package models

import (
	"github.com/globalsign/mgo/bson"
)

// structure of activity sensor data
type PowerSensorData struct {
	ID bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Timestamp TimeStamp `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	Watts float64 `json:"watts,omitempty" bson:"watts,omitempty"`
}
