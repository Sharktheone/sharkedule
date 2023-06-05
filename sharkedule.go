package main

import (
	"github.com/Sharktheone/sharkedule/api/router"
	"github.com/Sharktheone/sharkedule/database/db"
)

func main() {
	db.Init()
	router.Start()
}
