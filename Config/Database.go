package Config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"insight4wear-backend/Controllers"
	"log"
	"time"
)

// Database configurations for influxDB login

type DBConfig struct {
	Addr string
	DB string
	Username string
	Password string
}


func BuildMongoConnection()  {
	// Database config
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)

	// setup context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	// close connection when completed
	defer cancel()

	// ping our dbconnection
	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("Failed to establish connection to the database", err)
	} else {
		log.Println("Connected to the database!")
	}

	// create the database
	db := client.Database("insight4wear")

	// create the collections for the database TODO: do it for all sensors/data
	Controllers.ActivityCollection(db)
}
