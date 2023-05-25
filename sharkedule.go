package main

import (
	"sharkedule/api/router"
	"sharkedule/database"
)

func main() {
	database.Init()
	router.Start()
}
