package Models

import "go.mongodb.org/mongo-driver/bson/primitive"


// representing the shared values that each sensor will have
type CommonSensorData struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	StartTime string `json:"starttime,omitempty" bson:"starttime,omitempty"`
	EndTime string `json:"endtime,omitempty" bson:"endtime,omitempty"`
}

// structure of activity sensor data
type ActivitySensorData struct {
	CommonSensorData
	Activity int `json:"activity,omitempty" bson:"activity,omitempty"`
	Duration float32 `json:"duration,omitempty" bson:"duration,omitempty"`
}



