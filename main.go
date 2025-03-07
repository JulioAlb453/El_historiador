package main

import "main/routes"

func main() {
	router := routes.SetupRouter()

	router.Run(":8080") 
}