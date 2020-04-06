package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"insight4wear-backend/controllers"
	"log"
	"os"
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
	// Database db
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

	// ping our db connection
	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("Failed to establish connection to the database", err)
	} else {
		log.Println("Successfully connected to the database!")
	}

	// create the database
	db := client.Database(os.Getenv("DATABASE_NAME"))

	// create the collections for the database TODO: do it for all sensors/data
	controllers.ActivityCollection(db)
	controllers.CalorieCollection(db)
	controllers.HeartCollection(db)
	controllers.PowerCollection(db)
	controllers.StepCollection(db)
}
