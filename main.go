package main

import (
	"insight4wear-backend/Config"
	"insight4wear-backend/Routes"
)

func main()  {
	// start mongo configuration
	Config.BuildMongoConnection()

	// set up routes
	router := Routes.SetupRouter()

	router.Run()


}
