package main

import (
	"chal9/database"
	"chal9/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8000")
}
