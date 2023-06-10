package db

import (
	"github.com/Sharktheone/sharkedule/database"
	jsonfileV2 "github.com/Sharktheone/sharkedule/database/jsonfile"
	"log"
)

var (
	DB database.IDatabase
)

func Init() {
	DB = jsonfileV2.NewJSONFile()
	if err := DB.Load(); err != nil {
		log.Fatalf("failed loading database v2: %v", err)
	}
}
