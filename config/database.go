package config

import (

	"github.com/influxdata/influxdb1-client/v2"
	"log"
	"time"
)

// Database configurations for influxDB login
const (
	DB = "insight4wear"
	username = "userinsight"
	password = "insightpower"
)

// Function to connect to the database
func buildDBConfig() {
	connection, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:               "http://localhost:8086",
		Username:           username,
		Password:           password,
		UserAgent:          "",
		Timeout:            0,
		InsecureSkipVerify: false,
		TLSConfig:          nil,
		Proxy:              nil,
	})

	// check for errors
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	// creating a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Precision:        "ns",
		Database:         DB,
		RetentionPolicy:  "",
		WriteConsistency: "",
	})
	// check for errors creating batch for this database
	if err != nil {
		log.Fatal(err)
	}

	// create a point and add it to the database
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{} {
		"idle": 10.1,
		"system": 53.3,
		"user": 46.6,
	}

	pt, err := client.NewPoint("cpu_usage", tags, fields,time.Now())
	if err != nil {
		log.Fatal(err)
	}
	// add the point to the batch
	bp.AddPoint(pt)

	// write to the batch
	if err := connection.Write(bp); err != nil {
		log.Fatal(err)
	}

	// close the client resources
	if err := connection.Close(); err != nil {
		log.Fatal(err)
	}

}
