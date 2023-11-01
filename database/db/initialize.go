package db

import (
	"github.com/Sharktheone/sharkedule/database"
	"github.com/Sharktheone/sharkedule/database/jsonfile"
	"log"
)

var (
	DB database.IDatabase
)

func Init() {
	DB = jsonfile.NewJSONFile()
	if err := DB.Load(); err != nil {
		log.Fatalf("failed loading database v2: %v", err)
	}
}
