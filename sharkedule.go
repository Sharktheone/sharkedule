package main

import (
	"sharkedule/api/router"
	"sharkedule/database/db"
	"sharkedule/kanban"
)

func main() {
	db.Init()
	kanban.Load()
	router.Start()
}
