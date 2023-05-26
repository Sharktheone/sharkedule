package main

import (
	"sharkedule/api/router"
	"sharkedule/database/db"
)

func main() {
	db.Init()
	router.Start()
}
