package main

import "linkfi-coding/routes"

func main() {
	r := routes.SetupRoutes()
	r.Run(":8080")
}
