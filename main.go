package main

import (
	"lms/config/db"
	"lms/routes"
)

func main() {

	// Server to database
	db.Server()

	// Initalize the router
	routes.SetupRouter()

	// Run the server
	routes.SetupRouter().Run(":3000")
}
