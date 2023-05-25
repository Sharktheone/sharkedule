package main

import (
	"sharkedule/api/router"
	"sharkedule/database/dbInit"
)

func main() {
	dbInit.Init()
	router.Start()
}
