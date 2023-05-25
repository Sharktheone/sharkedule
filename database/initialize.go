package database

import (
	"log"
	"sharkedule/database/jsonfile"
)

var (
	DB IDatabase
)

func Init() {
	DB = jsonfile.NewJSONFile()
	if err := DB.Load(); err != nil {
		log.Fatalf("failed loading database: %v", err)
	}
}
