package main

import (
	"LATIHAN1/database"
	"LATIHAN1/routers"
)

func main() {
	database.StartDB()
	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}