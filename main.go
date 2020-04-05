package main

import (
	"github.com/joho/godotenv"
	"insight4wear-backend/db"
	"insight4wear-backend/routes"
	"log"
)

func init() {
	// Log error if there is no .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file exists")
	}
}

func main()  {
	// start mongo configuration
	db.BuildMongoConnection()

	// set up routes
	router := routes.SetupRouter()

	router.Run()


}
