package dbInit

import (
	"log"
	"sharkedule/database"
	"sharkedule/database/jsonfile"
)

var (
	DB database.IDatabase
)

func Init() {
	DB = jsonfile.NewJSONFile()
	if err := DB.Load(); err != nil {
		log.Fatalf("failed loading database: %v", err)
	}
}
