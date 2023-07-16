package main

import (
	"api-rest/database"
	routes "api-rest/routes"
	"fmt"
)

func main() {
	database.ConnectDatabase()
	fmt.Println("Starting Server")
	routes.HandleRequest()
}
