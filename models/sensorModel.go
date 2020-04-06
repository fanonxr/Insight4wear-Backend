package models

import "github.com/gin-gonic/gin"

// SensorData interface that each type of sensor model will implement to interact with the database
type SensorData interface {
	GetAllData(c *gin.Context)
	CreateData(c *gin.Context)
	DeleteData(c *gin.Context)
}

// representing the shared values that each sensor will have
type TimeStamp struct {
	StartTime string `json:"starttime,omitempty" bson:"starttime,omitempty"`
	EndTime string `json:"endtime,omitempty" bson:"endtime,omitempty"`
}